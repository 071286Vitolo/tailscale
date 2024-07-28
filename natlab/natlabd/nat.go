package main

import (
	"math/rand/v2"
	"net/netip"
	"time"
)

// NATTable is what a NAT implementation is expected to do.
//
// This project tests Tailscale as it faces various combinations various NAT
// implementations (e.g. Linux easy style NAT vs FreeBSD hard/endpoint dependent
// NAT vs Cloud 1:1 NAT, etc)
//
// Implementations of NATTable need not handle concurrency; the natlab serializes
// all calls into a NATTable.
//
// The provided `at` value will typically be time.Now, except for tests.
// Implementations should not use real time and should only compare
// previously provided time values.
type NATTable interface {
	// PickOutgoingSrc returns the source address to use for an outgoing packet.
	//
	// The result should either be invalid (to drop the packet) or a WAN (not
	// private) IP address.
	//
	// Typically, the src is a LAN source IP address, but it might also be a WAN
	// IP address if the packet is being forwarded for a source machine that has
	// a public IP address.
	PickOutgoingSrc(src, dst netip.AddrPort, at time.Time) (wanSrc netip.AddrPort)

	// PickIncomingDst returns the destination address to use for an incoming
	// packet. The incoming src address is always a public WAN IP.
	//
	// The result should either be invalid (to drop the packet) or the IP
	// address of a machine on the local network address, usually a private
	// LAN IP.
	PickIncomingDst(src, dst netip.AddrPort, at time.Time) (lanDst netip.AddrPort)
}

// oneToOneNAT is a 1:1 NAT, like a typical EC2 VM.
type oneToOneNAT struct {
	lanIP netip.Addr
	wanIP netip.Addr
}

var _ NATTable = (*oneToOneNAT)(nil)

func (n *oneToOneNAT) PickOutgoingSrc(src, dst netip.AddrPort, at time.Time) (wanSrc netip.AddrPort) {
	return netip.AddrPortFrom(n.wanIP, src.Port())
}

func (n *oneToOneNAT) PickIncomingDst(src, dst netip.AddrPort, at time.Time) (lanDst netip.AddrPort) {
	return netip.AddrPortFrom(n.lanIP, dst.Port())
}

type hardKeyOut struct {
	lanIP netip.Addr
	dst   netip.AddrPort
}

type hardKeyIn struct {
	wanPort uint16
	src     netip.AddrPort
}

type portMappingAndTime struct {
	port uint16
	at   time.Time
}

type lanAddrAndTime struct {
	lanAddr netip.AddrPort
	at      time.Time
}

// hardNAT is an "Endpoint Dependent" NAT, like FreeBSD/pfSense/OPNsense.
// This is shown as "MappingVariesByDestIP: true" by netcheck, and what
// Tailscale calls "Hard NAT".
type hardNAT struct {
	wanIP netip.Addr

	out map[hardKeyOut]portMappingAndTime
	in  map[hardKeyIn]lanAddrAndTime
}

var _ NATTable = (*hardNAT)(nil)

func (n *hardNAT) PickOutgoingSrc(src, dst netip.AddrPort, at time.Time) (wanSrc netip.AddrPort) {
	ko := hardKeyOut{src.Addr(), dst}
	if pm, ok := n.out[ko]; ok {
		// Existing flow.
		// TODO: bump timestamp
		return netip.AddrPortFrom(n.wanIP, pm.port)
	}

	// No existing mapping exists. Create one.

	// TODO: clean up old expired mappings

	// Instead of proper data structures that would be efficient, we instead
	// just loop a bunch and look for a free port. This project is only used
	// by tests and doesn't care about performance, this is good enough.
	port := src.Port() // start with the port the client used? (TODO: does FreeBSD do this? make knob?)
	for {
		ki := hardKeyIn{wanPort: port, src: dst}
		if _, ok := n.in[ki]; ok {
			// Collision. Try again.
			port = rand.N(uint16(32<<10)) + 32<<10 // pick some "ephemeral" port
			continue
		}
		n.in[ki] = lanAddrAndTime{lanAddr: src, at: at}
		n.out[ko] = portMappingAndTime{port: port, at: at}
	}
}

func (n *hardNAT) PickIncomingDst(src, dst netip.AddrPort, at time.Time) (lanDst netip.AddrPort) {
	if dst.Addr() != n.wanIP {
		return netip.AddrPort{} // drop; not for us. shouldn't happen if natlabd routing isn't broken.
	}
	ki := hardKeyIn{wanPort: dst.Port(), src: src}
	if pm, ok := n.in[ki]; ok {
		// Existing flow.
		return pm.lanAddr
	}
	return netip.AddrPort{} // drop; no mapping
}