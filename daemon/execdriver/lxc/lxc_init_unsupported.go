// +build !linux

package lxc

func setHostname(hostname string) error {
	panic("Not supported on darwin")
}
//
//func finalizeNamespace(args *execdriver.InitArgs) error {
//	panic("Not supported on darwin")
//}
