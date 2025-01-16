package frida

//#include <frida-core.h>
import "C"
import "unsafe"

// PortalMembership type is used to join portal with session.
type PortalMembership struct {
	mem *C.FridaPortalMembership
}

// Terminate terminates the session membership
func (p *PortalMembership) Terminate() error {
	var err *C.GError
	C.frida_portal_membership_terminate_sync(p.mem, nil, &err)
	return handleGError(err)
}

// Clean will clean the resources held by the portal membership.
func (p *PortalMembership) Clean() {
	clean(unsafe.Pointer(p.mem), unrefFrida)
}
