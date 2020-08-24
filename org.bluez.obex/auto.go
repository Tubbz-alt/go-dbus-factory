package obex

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "github.com/linuxdeepin/go-dbus-factory/object_manager"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type ObjectManager struct {
	object_manager.ObjectManager // interface org.freedesktop.DBus.ObjectManager
	proxy.Object
}

func NewObjectManager(conn *dbus.Conn) *ObjectManager {
	obj := new(ObjectManager)
	obj.Object.Init_(conn, "org.bluez.obex", "/")
	return obj
}

type Manager struct {
	agentManager // interface org.bluez.obex.AgentManager1
	client       // interface org.bluez.obex.Client1
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.bluez.obex", "/org/bluez/obex")
	return obj
}

func (obj *Manager) AgentManager() *agentManager {
	return &obj.agentManager
}

type agentManager struct{}

func (v *agentManager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*agentManager) GetInterfaceName_() string {
	return "org.bluez.obex.AgentManager1"
}

// method RegisterAgent

func (v *agentManager) GoRegisterAgent(flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RegisterAgent", flags, ch, agent)
}

func (v *agentManager) RegisterAgent(flags dbus.Flags, agent dbus.ObjectPath) error {
	return (<-v.GoRegisterAgent(flags, make(chan *dbus.Call, 1), agent).Done).Err
}

// method UnregisterAgent

func (v *agentManager) GoUnregisterAgent(flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnregisterAgent", flags, ch, agent)
}

func (v *agentManager) UnregisterAgent(flags dbus.Flags, agent dbus.ObjectPath) error {
	return (<-v.GoUnregisterAgent(flags, make(chan *dbus.Call, 1), agent).Done).Err
}

func (obj *Manager) Client() *client {
	return &obj.client
}

type client struct{}

func (v *client) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*client) GetInterfaceName_() string {
	return "org.bluez.obex.Client1"
}

// method CreateSession

func (v *client) GoCreateSession(flags dbus.Flags, ch chan *dbus.Call, destination string, args map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateSession", flags, ch, destination, args)
}

func (*client) StoreCreateSession(call *dbus.Call) (session dbus.ObjectPath, err error) {
	err = call.Store(&session)
	return
}

func (v *client) CreateSession(flags dbus.Flags, destination string, args map[string]dbus.Variant) (session dbus.ObjectPath, err error) {
	return v.StoreCreateSession(
		<-v.GoCreateSession(flags, make(chan *dbus.Call, 1), destination, args).Done)
}

// method RemoveSession

func (v *client) GoRemoveSession(flags dbus.Flags, ch chan *dbus.Call, session dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveSession", flags, ch, session)
}

func (v *client) RemoveSession(flags dbus.Flags, session dbus.ObjectPath) error {
	return (<-v.GoRemoveSession(flags, make(chan *dbus.Call, 1), session).Done).Err
}

type Session struct {
	session    // interface org.bluez.obex.Session1
	objectPush // interface org.bluez.obex.ObjectPush1
	proxy.Object
}

func NewSession(conn *dbus.Conn, path dbus.ObjectPath) (*Session, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Session)
	obj.Object.Init_(conn, "org.bluez.obex", path)
	return obj, nil
}

func (obj *Session) Session() *session {
	return &obj.session
}

type session struct{}

func (v *session) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*session) GetInterfaceName_() string {
	return "org.bluez.obex.Session1"
}

// method GetCapabilities

func (v *session) GoGetCapabilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCapabilities", flags, ch)
}

func (*session) StoreGetCapabilities(call *dbus.Call) (capabilities string, err error) {
	err = call.Store(&capabilities)
	return
}

func (v *session) GetCapabilities(flags dbus.Flags) (capabilities string, err error) {
	return v.StoreGetCapabilities(
		<-v.GoGetCapabilities(flags, make(chan *dbus.Call, 1)).Done)
}

// property Source s

func (v *session) Source() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Source",
	}
}

// property Destination s

func (v *session) Destination() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Destination",
	}
}

// property Channel y

func (v *session) Channel() proxy.PropByte {
	return proxy.PropByte{
		Impl: v,
		Name: "Channel",
	}
}

// property Target s

func (v *session) Target() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Target",
	}
}

func (obj *Session) ObjectPush() *objectPush {
	return &obj.objectPush
}

type objectPush struct{}

func (v *objectPush) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*objectPush) GetInterfaceName_() string {
	return "org.bluez.obex.ObjectPush1"
}

// method SendFile

func (v *objectPush) GoSendFile(flags dbus.Flags, ch chan *dbus.Call, sourcefile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SendFile", flags, ch, sourcefile)
}

func (*objectPush) StoreSendFile(call *dbus.Call) (transfer dbus.ObjectPath, properties map[string]dbus.Variant, err error) {
	err = call.Store(&transfer, &properties)
	return
}

func (v *objectPush) SendFile(flags dbus.Flags, sourcefile string) (transfer dbus.ObjectPath, properties map[string]dbus.Variant, err error) {
	return v.StoreSendFile(
		<-v.GoSendFile(flags, make(chan *dbus.Call, 1), sourcefile).Done)
}

// method PullBusinessCard

func (v *objectPush) GoPullBusinessCard(flags dbus.Flags, ch chan *dbus.Call, targetfile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PullBusinessCard", flags, ch, targetfile)
}

func (*objectPush) StorePullBusinessCard(call *dbus.Call) (transfer dbus.ObjectPath, properties map[string]dbus.Variant, err error) {
	err = call.Store(&transfer, &properties)
	return
}

func (v *objectPush) PullBusinessCard(flags dbus.Flags, targetfile string) (transfer dbus.ObjectPath, properties map[string]dbus.Variant, err error) {
	return v.StorePullBusinessCard(
		<-v.GoPullBusinessCard(flags, make(chan *dbus.Call, 1), targetfile).Done)
}

// method ExchangeBusinessCards

func (v *objectPush) GoExchangeBusinessCards(flags dbus.Flags, ch chan *dbus.Call, clientfile string, targetfile string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ExchangeBusinessCards", flags, ch, clientfile, targetfile)
}

func (*objectPush) StoreExchangeBusinessCards(call *dbus.Call) (transfer dbus.ObjectPath, properties map[string]dbus.Variant, err error) {
	err = call.Store(&transfer, &properties)
	return
}

func (v *objectPush) ExchangeBusinessCards(flags dbus.Flags, clientfile string, targetfile string) (transfer dbus.ObjectPath, properties map[string]dbus.Variant, err error) {
	return v.StoreExchangeBusinessCards(
		<-v.GoExchangeBusinessCards(flags, make(chan *dbus.Call, 1), clientfile, targetfile).Done)
}

type Transfer struct {
	transfer // interface org.bluez.obex.Transfer1
	proxy.Object
}

func NewTransfer(conn *dbus.Conn, path dbus.ObjectPath) (*Transfer, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Transfer)
	obj.Object.Init_(conn, "org.bluez.obex", path)
	return obj, nil
}

type transfer struct{}

func (v *transfer) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*transfer) GetInterfaceName_() string {
	return "org.bluez.obex.Transfer1"
}

// method Cancel

func (v *transfer) GoCancel(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Cancel", flags, ch)
}

func (v *transfer) Cancel(flags dbus.Flags) error {
	return (<-v.GoCancel(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property Status s

func (v *transfer) Status() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Status",
	}
}

// property Session o

func (v *transfer) Session() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Session",
	}
}

// property Name s

func (v *transfer) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Type s

func (v *transfer) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

// property Size t

func (v *transfer) Size() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Size",
	}
}

// property Time t

func (v *transfer) Time() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Time",
	}
}

// property Filename s

func (v *transfer) Filename() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Filename",
	}
}

// property Transferred t

func (v *transfer) Transferred() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Transferred",
	}
}
