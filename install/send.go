package install

import (
	"fmt"
	"path"
)

//SendPackage is
func (s *DmosInstaller) SendPackage() {
	pkg := path.Base(PkgUrl)
	// rm old Dmos in package avoid old version problem. if Dmos not exist in package then skip rm
	kubeHook := fmt.Sprintf("cd /root && rm -rf kube && tar zxvf %s  && cd /root/kube/shell && rm -f ../bin/Dmos && bash init.sh", pkg)
	deletekubectl := `sed -i '/kubectl/d;/Dmos/d' /root/.bashrc `
	completion := "echo 'command -v kubectl &>/dev/null && source <(kubectl completion bash)' >> /root/.bashrc && echo '[ -x /usr/bin/Dmos ] && source <(Dmos completion bash)' >> /root/.bashrc && source /root/.bashrc"
	kubeHook = kubeHook + " && " + deletekubectl + " && " + completion
	PkgUrl = SendPackage(PkgUrl, s.Hosts, "/root", nil, &kubeHook)

}

// SendDmos is send the exec Dmos to /usr/bin/Dmos
func (s *DmosInstaller) SendDmos() {
	// send Dmos first to avoid old version
	Dmos := FetchDmosAbsPath()
	beforeHook := "ps -ef |grep -v 'grep'|grep Dmos >/dev/null || rm -rf /usr/bin/Dmos"
	afterHook := "chmod a+x /usr/bin/Dmos"
	SendPackage(Dmos, s.Hosts, "/usr/bin", &beforeHook, &afterHook)
}

// SendPackage is send new pkg to all nodes.
func (u *DmosUpgrade) SendPackage() {
	all := append(u.Masters, u.Nodes...)
	pkg := path.Base(u.NewPkgUrl)
	// rm old Dmos in package avoid old version problem. if Dmos not exist in package then skip rm
	var kubeHook string
	if For120(Version) {
		// TODO update need load modprobe -- br_netfilter modprobe -- bridge.
		// https://github.com/fanux/cloud-kernel/issues/23
		kubeHook = fmt.Sprintf("cd /root && rm -rf kube && tar zxvf %s  && cd /root/kube/shell && rm -f ../bin/Dmos && (ctr -n=k8s.io image import ../images/images.tar || true) && cp -f ../bin/* /usr/bin/ ", pkg)
	} else {
		kubeHook = fmt.Sprintf("cd /root && rm -rf kube && tar zxvf %s  && cd /root/kube/shell && rm -f ../bin/Dmos && (docker load -i ../images/images.tar || true) && cp -f ../bin/* /usr/bin/ ", pkg)

	}

	PkgUrl = SendPackage(pkg, all, "/root", nil, &kubeHook)
}
