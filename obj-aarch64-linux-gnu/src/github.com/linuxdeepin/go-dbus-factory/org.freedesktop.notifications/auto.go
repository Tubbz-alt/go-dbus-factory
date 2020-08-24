package notifications

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Notifications struct {
	notifications // interface org.freedesktop.Notifications
	proxy.Object
}

func NewNotifications(conn *dbus.Conn) *Notifications {
	obj := new(Notifications)
	obj.Object.Init_(conn, "org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	return obj
}

type notifications struct{}

func (v *notifications) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*notifications) GetInterfaceName_() string {
	return "org.freedesktop.Notifications"
}

// method GetCapabilities

func (v *notifications) GoGetCapabilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCapabilities", flags, ch)
}

func (*notifications) StoreGetCapabilities(call *dbus.Call) (capabilities []string, err error) {
	err = call.Store(&capabilities)
	return
}

func (v *notifications) GetCapabilities(flags dbus.Flags) (capabilities []string, err error) {
	return v.StoreGetCapabilities(
		<-v.GoGetCapabilities(flags, make(chan *dbus.Call, 1)).Done)
}

// method Notify

func (v *notifications) GoNotify(flags dbus.Flags, ch chan *dbus.Call, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Notify", flags, ch, app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout)
}

func (*notifications) StoreNotify(call *dbus.Call) (id uint32, err error) {
	err = call.Store(&id)
	return
}

func (v *notifications) Notify(flags dbus.Flags, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) (id uint32, err error) {
	return v.StoreNotify(
		<-v.GoNotify(flags, make(chan *dbus.Call, 1), app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout).Done)
}

// method CloseNotification

func (v *notifications) GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CloseNotification", flags, ch, id)
}

func (v *notifications) CloseNotification(flags dbus.Flags, id uint32) error {
	return (<-v.GoCloseNotification(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// method GetServerInformation

func (v *notifications) GoGetServerInformation(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetServerInformation", flags, ch)
}

func (*notifications) StoreGetServerInformation(call *dbus.Call) (name string, vendor string, version string, spec_version string, err error) {
	err = call.Store(&name, &vendor, &version, &spec_version)
	return
}

func (v *notifications) GetServerInformation(flags dbus.Flags) (name string, vendor string, version string, spec_version string, err error) {
	return v.StoreGetServerInformation(
		<-v.GoGetServerInformation(flags, make(chan *dbus.Call, 1)).Done)
}

// signal NotificationClosed

func (v *notifications) ConnectNotificationClosed(cb func(id uint32, reason uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NotificationClosed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NotificationClosed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id uint32
		var reason uint32
		err := dbus.Store(sig.Body, &id, &reason)
		if err == nil {
			cb(id, reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ActionInvoked

func (v *notifications) ConnectActionInvoked(cb func(id uint32, action_key string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ActionInvoked", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ActionInvoked",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id uint32
		var action_key string
		err := dbus.Store(sig.Body, &id, &action_key)
		if err == nil {
			cb(id, action_key)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
