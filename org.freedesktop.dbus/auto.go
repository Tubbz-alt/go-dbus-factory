package dbus

import "errors"
import "fmt"
import "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type DBus struct {
	dbusIfc // interface org.freedesktop.DBus
	proxy.Object
}

func NewDBus(conn *dbus.Conn) *DBus {
	obj := new(DBus)
	obj.Object.Init_(conn, "org.freedesktop.DBus", "/org/freedesktop/DBus")
	return obj
}

type dbusIfc struct{}

func (v *dbusIfc) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*dbusIfc) GetInterfaceName_() string {
	return "org.freedesktop.DBus"
}

// method Hello

func (v *dbusIfc) GoHello(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hello", flags, ch)
}

func (*dbusIfc) StoreHello(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *dbusIfc) Hello(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreHello(
		<-v.GoHello(flags, make(chan *dbus.Call, 1)).Done)
}

// method RequestName

func (v *dbusIfc) GoRequestName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RequestName", flags, ch, arg0, arg1)
}

func (*dbusIfc) StoreRequestName(call *dbus.Call) (arg2 uint32, err error) {
	err = call.Store(&arg2)
	return
}

func (v *dbusIfc) RequestName(flags dbus.Flags, arg0 string, arg1 uint32) (arg2 uint32, err error) {
	return v.StoreRequestName(
		<-v.GoRequestName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method ReleaseName

func (v *dbusIfc) GoReleaseName(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseName", flags, ch, arg0)
}

func (*dbusIfc) StoreReleaseName(call *dbus.Call) (arg1 uint32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) ReleaseName(flags dbus.Flags, arg0 string) (arg1 uint32, err error) {
	return v.StoreReleaseName(
		<-v.GoReleaseName(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method StartServiceByName

func (v *dbusIfc) GoStartServiceByName(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartServiceByName", flags, ch, arg0, arg1)
}

func (*dbusIfc) StoreStartServiceByName(call *dbus.Call) (arg2 uint32, err error) {
	err = call.Store(&arg2)
	return
}

func (v *dbusIfc) StartServiceByName(flags dbus.Flags, arg0 string, arg1 uint32) (arg2 uint32, err error) {
	return v.StoreStartServiceByName(
		<-v.GoStartServiceByName(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method UpdateActivationEnvironment

func (v *dbusIfc) GoUpdateActivationEnvironment(flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateActivationEnvironment", flags, ch, arg0)
}

func (v *dbusIfc) UpdateActivationEnvironment(flags dbus.Flags, arg0 map[string]string) error {
	return (<-v.GoUpdateActivationEnvironment(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method NameHasOwner

func (v *dbusIfc) GoNameHasOwner(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NameHasOwner", flags, ch, arg0)
}

func (*dbusIfc) StoreNameHasOwner(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) NameHasOwner(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StoreNameHasOwner(
		<-v.GoNameHasOwner(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListNames

func (v *dbusIfc) GoListNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListNames", flags, ch)
}

func (*dbusIfc) StoreListNames(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *dbusIfc) ListNames(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreListNames(
		<-v.GoListNames(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListActivatableNames

func (v *dbusIfc) GoListActivatableNames(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListActivatableNames", flags, ch)
}

func (*dbusIfc) StoreListActivatableNames(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *dbusIfc) ListActivatableNames(flags dbus.Flags) (arg0 []string, err error) {
	return v.StoreListActivatableNames(
		<-v.GoListActivatableNames(flags, make(chan *dbus.Call, 1)).Done)
}

// method AddMatch

func (v *dbusIfc) GoAddMatch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddMatch", flags, ch, arg0)
}

func (v *dbusIfc) AddMatch(flags dbus.Flags, arg0 string) error {
	return (<-v.GoAddMatch(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method RemoveMatch

func (v *dbusIfc) GoRemoveMatch(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveMatch", flags, ch, arg0)
}

func (v *dbusIfc) RemoveMatch(flags dbus.Flags, arg0 string) error {
	return (<-v.GoRemoveMatch(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method GetNameOwner

func (v *dbusIfc) GoGetNameOwner(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetNameOwner", flags, ch, arg0)
}

func (*dbusIfc) StoreGetNameOwner(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetNameOwner(flags dbus.Flags, arg0 string) (arg1 string, err error) {
	return v.StoreGetNameOwner(
		<-v.GoGetNameOwner(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListQueuedOwners

func (v *dbusIfc) GoListQueuedOwners(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListQueuedOwners", flags, ch, arg0)
}

func (*dbusIfc) StoreListQueuedOwners(call *dbus.Call) (arg1 []string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) ListQueuedOwners(flags dbus.Flags, arg0 string) (arg1 []string, err error) {
	return v.StoreListQueuedOwners(
		<-v.GoListQueuedOwners(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetConnectionUnixUser

func (v *dbusIfc) GoGetConnectionUnixUser(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionUnixUser", flags, ch, arg0)
}

func (*dbusIfc) StoreGetConnectionUnixUser(call *dbus.Call) (arg1 uint32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetConnectionUnixUser(flags dbus.Flags, arg0 string) (arg1 uint32, err error) {
	return v.StoreGetConnectionUnixUser(
		<-v.GoGetConnectionUnixUser(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetConnectionUnixProcessID

func (v *dbusIfc) GoGetConnectionUnixProcessID(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionUnixProcessID", flags, ch, arg0)
}

func (*dbusIfc) StoreGetConnectionUnixProcessID(call *dbus.Call) (arg1 uint32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetConnectionUnixProcessID(flags dbus.Flags, arg0 string) (arg1 uint32, err error) {
	return v.StoreGetConnectionUnixProcessID(
		<-v.GoGetConnectionUnixProcessID(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetAdtAuditSessionData

func (v *dbusIfc) GoGetAdtAuditSessionData(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAdtAuditSessionData", flags, ch, arg0)
}

func (*dbusIfc) StoreGetAdtAuditSessionData(call *dbus.Call) (arg1 []uint8, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetAdtAuditSessionData(flags dbus.Flags, arg0 string) (arg1 []uint8, err error) {
	return v.StoreGetAdtAuditSessionData(
		<-v.GoGetAdtAuditSessionData(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetConnectionSELinuxSecurityContext

func (v *dbusIfc) GoGetConnectionSELinuxSecurityContext(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionSELinuxSecurityContext", flags, ch, arg0)
}

func (*dbusIfc) StoreGetConnectionSELinuxSecurityContext(call *dbus.Call) (arg1 []uint8, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetConnectionSELinuxSecurityContext(flags dbus.Flags, arg0 string) (arg1 []uint8, err error) {
	return v.StoreGetConnectionSELinuxSecurityContext(
		<-v.GoGetConnectionSELinuxSecurityContext(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ReloadConfig

func (v *dbusIfc) GoReloadConfig(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReloadConfig", flags, ch)
}

func (v *dbusIfc) ReloadConfig(flags dbus.Flags) error {
	return (<-v.GoReloadConfig(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetId

func (v *dbusIfc) GoGetId(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetId", flags, ch)
}

func (*dbusIfc) StoreGetId(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *dbusIfc) GetId(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreGetId(
		<-v.GoGetId(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetConnectionCredentials

func (v *dbusIfc) GoGetConnectionCredentials(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetConnectionCredentials", flags, ch, arg0)
}

func (*dbusIfc) StoreGetConnectionCredentials(call *dbus.Call) (arg1 map[string]dbus.Variant, err error) {
	err = call.Store(&arg1)
	return
}

func (v *dbusIfc) GetConnectionCredentials(flags dbus.Flags, arg0 string) (arg1 map[string]dbus.Variant, err error) {
	return v.StoreGetConnectionCredentials(
		<-v.GoGetConnectionCredentials(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// signal NameOwnerChanged

func (v *dbusIfc) ConnectNameOwnerChanged(cb func(arg0 string, arg1 string, arg2 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameOwnerChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameOwnerChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 string
		var arg2 string
		err := dbus.Store(sig.Body, &arg0, &arg1, &arg2)
		if err == nil {
			cb(arg0, arg1, arg2)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NameLost

func (v *dbusIfc) ConnectNameLost(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameLost", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameLost",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal NameAcquired

func (v *dbusIfc) ConnectNameAcquired(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NameAcquired", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NameAcquired",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
