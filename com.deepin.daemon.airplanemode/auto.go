package airplanemode

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

type AirplaneMode struct {
	airplaneMode // interface com.deepin.daemon.AirplaneMode
	proxy.Object
}

func NewAirplaneMode(conn *dbus.Conn) *AirplaneMode {
	obj := new(AirplaneMode)
	obj.Object.Init_(conn, "com.deepin.daemon.AirplaneMode", "/com/deepin/daemon/AirplaneMode")
	return obj
}

type airplaneMode struct{}

func (v *airplaneMode) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*airplaneMode) GetInterfaceName_() string {
	return "com.deepin.daemon.AirplaneMode"
}

// method DumpState

func (v *airplaneMode) GoDumpState(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DumpState", flags, ch)
}

func (v *airplaneMode) DumpState(flags dbus.Flags) error {
	return (<-v.GoDumpState(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Enable

func (v *airplaneMode) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, enabled)
}

func (v *airplaneMode) Enable(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method EnableBluetooth

func (v *airplaneMode) GoEnableBluetooth(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableBluetooth", flags, ch, enabled)
}

func (v *airplaneMode) EnableBluetooth(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableBluetooth(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method EnableWifi

func (v *airplaneMode) GoEnableWifi(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableWifi", flags, ch, enabled)
}

func (v *airplaneMode) EnableWifi(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableWifi(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// property Enabled b

func (v *airplaneMode) Enabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Enabled",
	}
}

// property WifiEnabled b

func (v *airplaneMode) WifiEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WifiEnabled",
	}
}

// property BluetoothEnabled b

func (v *airplaneMode) BluetoothEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "BluetoothEnabled",
	}
}
