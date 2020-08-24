package login1

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

type Manager struct {
	manager // interface org.freedesktop.login1.Manager
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.freedesktop.login1", "/org/freedesktop/login1")
	return obj
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "org.freedesktop.login1.Manager"
}

// method GetSession

func (v *manager) GoGetSession(flags dbus.Flags, ch chan *dbus.Call, sessionId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSession", flags, ch, sessionId)
}

func (*manager) StoreGetSession(call *dbus.Call) (sessionPath dbus.ObjectPath, err error) {
	err = call.Store(&sessionPath)
	return
}

func (v *manager) GetSession(flags dbus.Flags, sessionId string) (sessionPath dbus.ObjectPath, err error) {
	return v.StoreGetSession(
		<-v.GoGetSession(flags, make(chan *dbus.Call, 1), sessionId).Done)
}

// method GetSessionByPID

func (v *manager) GoGetSessionByPID(flags dbus.Flags, ch chan *dbus.Call, pid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSessionByPID", flags, ch, pid)
}

func (*manager) StoreGetSessionByPID(call *dbus.Call) (sessionPath dbus.ObjectPath, err error) {
	err = call.Store(&sessionPath)
	return
}

func (v *manager) GetSessionByPID(flags dbus.Flags, pid uint32) (sessionPath dbus.ObjectPath, err error) {
	return v.StoreGetSessionByPID(
		<-v.GoGetSessionByPID(flags, make(chan *dbus.Call, 1), pid).Done)
}

// method GetUser

func (v *manager) GoGetUser(flags dbus.Flags, ch chan *dbus.Call, uid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUser", flags, ch, uid)
}

func (*manager) StoreGetUser(call *dbus.Call) (userPath dbus.ObjectPath, err error) {
	err = call.Store(&userPath)
	return
}

func (v *manager) GetUser(flags dbus.Flags, uid uint32) (userPath dbus.ObjectPath, err error) {
	return v.StoreGetUser(
		<-v.GoGetUser(flags, make(chan *dbus.Call, 1), uid).Done)
}

// method GetUserByPID

func (v *manager) GoGetUserByPID(flags dbus.Flags, ch chan *dbus.Call, pid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetUserByPID", flags, ch, pid)
}

func (*manager) StoreGetUserByPID(call *dbus.Call) (userPath dbus.ObjectPath, err error) {
	err = call.Store(&userPath)
	return
}

func (v *manager) GetUserByPID(flags dbus.Flags, pid uint32) (userPath dbus.ObjectPath, err error) {
	return v.StoreGetUserByPID(
		<-v.GoGetUserByPID(flags, make(chan *dbus.Call, 1), pid).Done)
}

// method GetSeat

func (v *manager) GoGetSeat(flags dbus.Flags, ch chan *dbus.Call, seatId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSeat", flags, ch, seatId)
}

func (*manager) StoreGetSeat(call *dbus.Call) (seatPath dbus.ObjectPath, err error) {
	err = call.Store(&seatPath)
	return
}

func (v *manager) GetSeat(flags dbus.Flags, seatId string) (seatPath dbus.ObjectPath, err error) {
	return v.StoreGetSeat(
		<-v.GoGetSeat(flags, make(chan *dbus.Call, 1), seatId).Done)
}

// method ListSessions

func (v *manager) GoListSessions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListSessions", flags, ch)
}

func (*manager) StoreListSessions(call *dbus.Call) (sessionList []SessionDetail, err error) {
	err = call.Store(&sessionList)
	return
}

func (v *manager) ListSessions(flags dbus.Flags) (sessionList []SessionDetail, err error) {
	return v.StoreListSessions(
		<-v.GoListSessions(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListUsers

func (v *manager) GoListUsers(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListUsers", flags, ch)
}

func (*manager) StoreListUsers(call *dbus.Call) (userList []UserDetail, err error) {
	err = call.Store(&userList)
	return
}

func (v *manager) ListUsers(flags dbus.Flags) (userList []UserDetail, err error) {
	return v.StoreListUsers(
		<-v.GoListUsers(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListSeats

func (v *manager) GoListSeats(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListSeats", flags, ch)
}

func (*manager) StoreListSeats(call *dbus.Call) (seatList []SeatInfo, err error) {
	err = call.Store(&seatList)
	return
}

func (v *manager) ListSeats(flags dbus.Flags) (seatList []SeatInfo, err error) {
	return v.StoreListSeats(
		<-v.GoListSeats(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListInhibitors

func (v *manager) GoListInhibitors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListInhibitors", flags, ch)
}

func (*manager) StoreListInhibitors(call *dbus.Call) (inhibitorList []InhibitorInfo, err error) {
	err = call.Store(&inhibitorList)
	return
}

func (v *manager) ListInhibitors(flags dbus.Flags) (inhibitorList []InhibitorInfo, err error) {
	return v.StoreListInhibitors(
		<-v.GoListInhibitors(flags, make(chan *dbus.Call, 1)).Done)
}

// method CreateSession

func (v *manager) GoCreateSession(flags dbus.Flags, ch chan *dbus.Call, uid uint32, pid uint32, service string, type0 string, class string, desktop string, seatId string, vtnr uint32, tty string, display string, remote bool, remoteUser string, remoteHost string, properties [][]interface{}) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateSession", flags, ch, uid, pid, service, type0, class, desktop, seatId, vtnr, tty, display, remote, remoteUser, remoteHost, properties)
}

func (*manager) StoreCreateSession(call *dbus.Call) (sessionId string, sessionPath dbus.ObjectPath, runtimePath string, fifoFd dbus.UnixFD, uid0 uint32, seatId0 string, vtnr0 uint32, existing bool, err error) {
	err = call.Store(&sessionId, &sessionPath, &runtimePath, &fifoFd, &uid0, &seatId0, &vtnr0, &existing)
	return
}

func (v *manager) CreateSession(flags dbus.Flags, uid uint32, pid uint32, service string, type0 string, class string, desktop string, seatId string, vtnr uint32, tty string, display string, remote bool, remoteUser string, remoteHost string, properties [][]interface{}) (sessionId string, sessionPath dbus.ObjectPath, runtimePath string, fifoFd dbus.UnixFD, uid0 uint32, seatId0 string, vtnr0 uint32, existing bool, err error) {
	return v.StoreCreateSession(
		<-v.GoCreateSession(flags, make(chan *dbus.Call, 1), uid, pid, service, type0, class, desktop, seatId, vtnr, tty, display, remote, remoteUser, remoteHost, properties).Done)
}

// method ReleaseSession

func (v *manager) GoReleaseSession(flags dbus.Flags, ch chan *dbus.Call, sessionId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseSession", flags, ch, sessionId)
}

func (v *manager) ReleaseSession(flags dbus.Flags, sessionId string) error {
	return (<-v.GoReleaseSession(flags, make(chan *dbus.Call, 1), sessionId).Done).Err
}

// method ActivateSession

func (v *manager) GoActivateSession(flags dbus.Flags, ch chan *dbus.Call, sessionId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateSession", flags, ch, sessionId)
}

func (v *manager) ActivateSession(flags dbus.Flags, sessionId string) error {
	return (<-v.GoActivateSession(flags, make(chan *dbus.Call, 1), sessionId).Done).Err
}

// method ActivateSessionOnSeat

func (v *manager) GoActivateSessionOnSeat(flags dbus.Flags, ch chan *dbus.Call, sessionId string, seatId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateSessionOnSeat", flags, ch, sessionId, seatId)
}

func (v *manager) ActivateSessionOnSeat(flags dbus.Flags, sessionId string, seatId string) error {
	return (<-v.GoActivateSessionOnSeat(flags, make(chan *dbus.Call, 1), sessionId, seatId).Done).Err
}

// method LockSession

func (v *manager) GoLockSession(flags dbus.Flags, ch chan *dbus.Call, sessionId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LockSession", flags, ch, sessionId)
}

func (v *manager) LockSession(flags dbus.Flags, sessionId string) error {
	return (<-v.GoLockSession(flags, make(chan *dbus.Call, 1), sessionId).Done).Err
}

// method UnlockSession

func (v *manager) GoUnlockSession(flags dbus.Flags, ch chan *dbus.Call, sessionId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnlockSession", flags, ch, sessionId)
}

func (v *manager) UnlockSession(flags dbus.Flags, sessionId string) error {
	return (<-v.GoUnlockSession(flags, make(chan *dbus.Call, 1), sessionId).Done).Err
}

// method LockSessions

func (v *manager) GoLockSessions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LockSessions", flags, ch)
}

func (v *manager) LockSessions(flags dbus.Flags) error {
	return (<-v.GoLockSessions(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method UnlockSessions

func (v *manager) GoUnlockSessions(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnlockSessions", flags, ch)
}

func (v *manager) UnlockSessions(flags dbus.Flags) error {
	return (<-v.GoUnlockSessions(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method KillSession

func (v *manager) GoKillSession(flags dbus.Flags, ch chan *dbus.Call, sessionId string, who string, signo int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".KillSession", flags, ch, sessionId, who, signo)
}

func (v *manager) KillSession(flags dbus.Flags, sessionId string, who string, signo int32) error {
	return (<-v.GoKillSession(flags, make(chan *dbus.Call, 1), sessionId, who, signo).Done).Err
}

// method KillUser

func (v *manager) GoKillUser(flags dbus.Flags, ch chan *dbus.Call, uid uint32, signo int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".KillUser", flags, ch, uid, signo)
}

func (v *manager) KillUser(flags dbus.Flags, uid uint32, signo int32) error {
	return (<-v.GoKillUser(flags, make(chan *dbus.Call, 1), uid, signo).Done).Err
}

// method TerminateSession

func (v *manager) GoTerminateSession(flags dbus.Flags, ch chan *dbus.Call, sessionId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TerminateSession", flags, ch, sessionId)
}

func (v *manager) TerminateSession(flags dbus.Flags, sessionId string) error {
	return (<-v.GoTerminateSession(flags, make(chan *dbus.Call, 1), sessionId).Done).Err
}

// method TerminateUser

func (v *manager) GoTerminateUser(flags dbus.Flags, ch chan *dbus.Call, uid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TerminateUser", flags, ch, uid)
}

func (v *manager) TerminateUser(flags dbus.Flags, uid uint32) error {
	return (<-v.GoTerminateUser(flags, make(chan *dbus.Call, 1), uid).Done).Err
}

// method TerminateSeat

func (v *manager) GoTerminateSeat(flags dbus.Flags, ch chan *dbus.Call, seatId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TerminateSeat", flags, ch, seatId)
}

func (v *manager) TerminateSeat(flags dbus.Flags, seatId string) error {
	return (<-v.GoTerminateSeat(flags, make(chan *dbus.Call, 1), seatId).Done).Err
}

// method SetUserLinger

func (v *manager) GoSetUserLinger(flags dbus.Flags, ch chan *dbus.Call, uid uint32, linger bool, interactive bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetUserLinger", flags, ch, uid, linger, interactive)
}

func (v *manager) SetUserLinger(flags dbus.Flags, uid uint32, linger bool, interactive bool) error {
	return (<-v.GoSetUserLinger(flags, make(chan *dbus.Call, 1), uid, linger, interactive).Done).Err
}

// method AttachDevice

func (v *manager) GoAttachDevice(flags dbus.Flags, ch chan *dbus.Call, seatId string, sysfs string, interactive bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AttachDevice", flags, ch, seatId, sysfs, interactive)
}

func (v *manager) AttachDevice(flags dbus.Flags, seatId string, sysfs string, interactive bool) error {
	return (<-v.GoAttachDevice(flags, make(chan *dbus.Call, 1), seatId, sysfs, interactive).Done).Err
}

// method FlushDevices

func (v *manager) GoFlushDevices(flags dbus.Flags, ch chan *dbus.Call, interactive bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FlushDevices", flags, ch, interactive)
}

func (v *manager) FlushDevices(flags dbus.Flags, interactive bool) error {
	return (<-v.GoFlushDevices(flags, make(chan *dbus.Call, 1), interactive).Done).Err
}

// method PowerOff

func (v *manager) GoPowerOff(flags dbus.Flags, ch chan *dbus.Call, interactive bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PowerOff", flags, ch, interactive)
}

func (v *manager) PowerOff(flags dbus.Flags, interactive bool) error {
	return (<-v.GoPowerOff(flags, make(chan *dbus.Call, 1), interactive).Done).Err
}

// method Reboot

func (v *manager) GoReboot(flags dbus.Flags, ch chan *dbus.Call, interactive bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reboot", flags, ch, interactive)
}

func (v *manager) Reboot(flags dbus.Flags, interactive bool) error {
	return (<-v.GoReboot(flags, make(chan *dbus.Call, 1), interactive).Done).Err
}

// method Suspend

func (v *manager) GoSuspend(flags dbus.Flags, ch chan *dbus.Call, interactive bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Suspend", flags, ch, interactive)
}

func (v *manager) Suspend(flags dbus.Flags, interactive bool) error {
	return (<-v.GoSuspend(flags, make(chan *dbus.Call, 1), interactive).Done).Err
}

// method Hibernate

func (v *manager) GoHibernate(flags dbus.Flags, ch chan *dbus.Call, interactive bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hibernate", flags, ch, interactive)
}

func (v *manager) Hibernate(flags dbus.Flags, interactive bool) error {
	return (<-v.GoHibernate(flags, make(chan *dbus.Call, 1), interactive).Done).Err
}

// method HybridSleep

func (v *manager) GoHybridSleep(flags dbus.Flags, ch chan *dbus.Call, interactive bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".HybridSleep", flags, ch, interactive)
}

func (v *manager) HybridSleep(flags dbus.Flags, interactive bool) error {
	return (<-v.GoHybridSleep(flags, make(chan *dbus.Call, 1), interactive).Done).Err
}

// method CanPowerOff

func (v *manager) GoCanPowerOff(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanPowerOff", flags, ch)
}

func (*manager) StoreCanPowerOff(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *manager) CanPowerOff(flags dbus.Flags) (result string, err error) {
	return v.StoreCanPowerOff(
		<-v.GoCanPowerOff(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanReboot

func (v *manager) GoCanReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanReboot", flags, ch)
}

func (*manager) StoreCanReboot(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *manager) CanReboot(flags dbus.Flags) (result string, err error) {
	return v.StoreCanReboot(
		<-v.GoCanReboot(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanSuspend

func (v *manager) GoCanSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanSuspend", flags, ch)
}

func (*manager) StoreCanSuspend(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *manager) CanSuspend(flags dbus.Flags) (result string, err error) {
	return v.StoreCanSuspend(
		<-v.GoCanSuspend(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanHibernate

func (v *manager) GoCanHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanHibernate", flags, ch)
}

func (*manager) StoreCanHibernate(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *manager) CanHibernate(flags dbus.Flags) (result string, err error) {
	return v.StoreCanHibernate(
		<-v.GoCanHibernate(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanHybridSleep

func (v *manager) GoCanHybridSleep(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanHybridSleep", flags, ch)
}

func (*manager) StoreCanHybridSleep(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *manager) CanHybridSleep(flags dbus.Flags) (result string, err error) {
	return v.StoreCanHybridSleep(
		<-v.GoCanHybridSleep(flags, make(chan *dbus.Call, 1)).Done)
}

// method ScheduleShutdown

func (v *manager) GoScheduleShutdown(flags dbus.Flags, ch chan *dbus.Call, type0 string, usec uint64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ScheduleShutdown", flags, ch, type0, usec)
}

func (v *manager) ScheduleShutdown(flags dbus.Flags, type0 string, usec uint64) error {
	return (<-v.GoScheduleShutdown(flags, make(chan *dbus.Call, 1), type0, usec).Done).Err
}

// method CancelScheduledShutdown

func (v *manager) GoCancelScheduledShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelScheduledShutdown", flags, ch)
}

func (*manager) StoreCancelScheduledShutdown(call *dbus.Call) (cancelled bool, err error) {
	err = call.Store(&cancelled)
	return
}

func (v *manager) CancelScheduledShutdown(flags dbus.Flags) (cancelled bool, err error) {
	return v.StoreCancelScheduledShutdown(
		<-v.GoCancelScheduledShutdown(flags, make(chan *dbus.Call, 1)).Done)
}

// method Inhibit

func (v *manager) GoInhibit(flags dbus.Flags, ch chan *dbus.Call, what string, who string, why string, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Inhibit", flags, ch, what, who, why, mode)
}

func (*manager) StoreInhibit(call *dbus.Call) (pipeFd dbus.UnixFD, err error) {
	err = call.Store(&pipeFd)
	return
}

func (v *manager) Inhibit(flags dbus.Flags, what string, who string, why string, mode string) (pipeFd dbus.UnixFD, err error) {
	return v.StoreInhibit(
		<-v.GoInhibit(flags, make(chan *dbus.Call, 1), what, who, why, mode).Done)
}

// method CanRebootToFirmwareSetup

func (v *manager) GoCanRebootToFirmwareSetup(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanRebootToFirmwareSetup", flags, ch)
}

func (*manager) StoreCanRebootToFirmwareSetup(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *manager) CanRebootToFirmwareSetup(flags dbus.Flags) (result string, err error) {
	return v.StoreCanRebootToFirmwareSetup(
		<-v.GoCanRebootToFirmwareSetup(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetRebootToFirmwareSetup

func (v *manager) GoSetRebootToFirmwareSetup(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRebootToFirmwareSetup", flags, ch, enable)
}

func (v *manager) SetRebootToFirmwareSetup(flags dbus.Flags, enable bool) error {
	return (<-v.GoSetRebootToFirmwareSetup(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// method SetWallMessage

func (v *manager) GoSetWallMessage(flags dbus.Flags, ch chan *dbus.Call, wallMessage string, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWallMessage", flags, ch, wallMessage, enable)
}

func (v *manager) SetWallMessage(flags dbus.Flags, wallMessage string, enable bool) error {
	return (<-v.GoSetWallMessage(flags, make(chan *dbus.Call, 1), wallMessage, enable).Done).Err
}

// signal SessionNew

func (v *manager) ConnectSessionNew(cb func(sessionId string, sessionPath dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SessionNew", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SessionNew",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var sessionId string
		var sessionPath dbus.ObjectPath
		err := dbus.Store(sig.Body, &sessionId, &sessionPath)
		if err == nil {
			cb(sessionId, sessionPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SessionRemoved

func (v *manager) ConnectSessionRemoved(cb func(sessionId string, sessionPath dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SessionRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SessionRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var sessionId string
		var sessionPath dbus.ObjectPath
		err := dbus.Store(sig.Body, &sessionId, &sessionPath)
		if err == nil {
			cb(sessionId, sessionPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UserNew

func (v *manager) ConnectUserNew(cb func(uid uint32, userPath dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UserNew", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UserNew",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var uid uint32
		var userPath dbus.ObjectPath
		err := dbus.Store(sig.Body, &uid, &userPath)
		if err == nil {
			cb(uid, userPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UserRemoved

func (v *manager) ConnectUserRemoved(cb func(uid uint32, userPath dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UserRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UserRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var uid uint32
		var userPath dbus.ObjectPath
		err := dbus.Store(sig.Body, &uid, &userPath)
		if err == nil {
			cb(uid, userPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SeatNew

func (v *manager) ConnectSeatNew(cb func(seatId string, seatPath dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SeatNew", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SeatNew",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var seatId string
		var seatPath dbus.ObjectPath
		err := dbus.Store(sig.Body, &seatId, &seatPath)
		if err == nil {
			cb(seatId, seatPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SeatRemoved

func (v *manager) ConnectSeatRemoved(cb func(seatId string, seatPath dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SeatRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SeatRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var seatId string
		var seatPath dbus.ObjectPath
		err := dbus.Store(sig.Body, &seatId, &seatPath)
		if err == nil {
			cb(seatId, seatPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PrepareForShutdown

func (v *manager) ConnectPrepareForShutdown(cb func(start bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PrepareForShutdown", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PrepareForShutdown",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var start bool
		err := dbus.Store(sig.Body, &start)
		if err == nil {
			cb(start)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PrepareForSleep

func (v *manager) ConnectPrepareForSleep(cb func(start bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PrepareForSleep", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PrepareForSleep",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var start bool
		err := dbus.Store(sig.Body, &start)
		if err == nil {
			cb(start)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property EnableWallMessages b

func (v *manager) EnableWallMessages() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "EnableWallMessages",
	}
}

// property WallMessage s

func (v *manager) WallMessage() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WallMessage",
	}
}

// property NAutoVTs u

func (v *manager) NAutoVTs() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "NAutoVTs",
	}
}

// property KillOnlyUsers as

func (v *manager) KillOnlyUsers() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "KillOnlyUsers",
	}
}

// property KillExcludeUsers as

func (v *manager) KillExcludeUsers() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "KillExcludeUsers",
	}
}

// property KillUserProcesses b

func (v *manager) KillUserProcesses() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "KillUserProcesses",
	}
}

// property RebootToFirmwareSetup b

func (v *manager) RebootToFirmwareSetup() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "RebootToFirmwareSetup",
	}
}

// property IdleHint b

func (v *manager) IdleHint() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IdleHint",
	}
}

// property IdleSinceHint t

func (v *manager) IdleSinceHint() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IdleSinceHint",
	}
}

// property IdleSinceHintMonotonic t

func (v *manager) IdleSinceHintMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IdleSinceHintMonotonic",
	}
}

// property BlockInhibited s

func (v *manager) BlockInhibited() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "BlockInhibited",
	}
}

// property DelayInhibited s

func (v *manager) DelayInhibited() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DelayInhibited",
	}
}

// property InhibitDelayMaxUSec t

func (v *manager) InhibitDelayMaxUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InhibitDelayMaxUSec",
	}
}

// property HandlePowerKey s

func (v *manager) HandlePowerKey() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HandlePowerKey",
	}
}

// property HandleSuspendKey s

func (v *manager) HandleSuspendKey() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HandleSuspendKey",
	}
}

// property HandleHibernateKey s

func (v *manager) HandleHibernateKey() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HandleHibernateKey",
	}
}

// property HandleLidSwitch s

func (v *manager) HandleLidSwitch() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HandleLidSwitch",
	}
}

// property HandleLidSwitchDocked s

func (v *manager) HandleLidSwitchDocked() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HandleLidSwitchDocked",
	}
}

// property HoldoffTimeoutUSec t

func (v *manager) HoldoffTimeoutUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "HoldoffTimeoutUSec",
	}
}

// property IdleAction s

func (v *manager) IdleAction() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IdleAction",
	}
}

// property IdleActionUSec t

func (v *manager) IdleActionUSec() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IdleActionUSec",
	}
}

// property PreparingForShutdown b

func (v *manager) PreparingForShutdown() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PreparingForShutdown",
	}
}

// property PreparingForSleep b

func (v *manager) PreparingForSleep() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PreparingForSleep",
	}
}

// property ScheduledShutdown (st)

func (v *manager) ScheduledShutdown() PropManagerScheduledShutdown {
	return PropManagerScheduledShutdown{
		Impl: v,
	}
}

type PropManagerScheduledShutdown struct {
	Impl proxy.Implementer
}

func (p PropManagerScheduledShutdown) Get(flags dbus.Flags) (value ScheduledShutdown, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"ScheduledShutdown", &value)
	return
}

func (p PropManagerScheduledShutdown) ConnectChanged(cb func(hasValue bool, value ScheduledShutdown)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v ScheduledShutdown
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, ScheduledShutdown{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"ScheduledShutdown", cb0)
}

// property Docked b

func (v *manager) Docked() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Docked",
	}
}

// property RemoveIPC b

func (v *manager) RemoveIPC() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "RemoveIPC",
	}
}

// property RuntimeDirectorySize t

func (v *manager) RuntimeDirectorySize() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "RuntimeDirectorySize",
	}
}

// property InhibitorsMax t

func (v *manager) InhibitorsMax() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "InhibitorsMax",
	}
}

// property NCurrentInhibitors t

func (v *manager) NCurrentInhibitors() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "NCurrentInhibitors",
	}
}

// property SessionsMax t

func (v *manager) SessionsMax() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "SessionsMax",
	}
}

// property NCurrentSessions t

func (v *manager) NCurrentSessions() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "NCurrentSessions",
	}
}

// property UserTasksMax t

func (v *manager) UserTasksMax() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "UserTasksMax",
	}
}

type Seat struct {
	seat // interface org.freedesktop.login1.Seat
	proxy.Object
}

func NewSeat(conn *dbus.Conn, path dbus.ObjectPath) (*Seat, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Seat)
	obj.Object.Init_(conn, "org.freedesktop.login1", path)
	return obj, nil
}

type seat struct{}

func (v *seat) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*seat) GetInterfaceName_() string {
	return "org.freedesktop.login1.Seat"
}

// method Terminate

func (v *seat) GoTerminate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Terminate", flags, ch)
}

func (v *seat) Terminate(flags dbus.Flags) error {
	return (<-v.GoTerminate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ActivateSession

func (v *seat) GoActivateSession(flags dbus.Flags, ch chan *dbus.Call, sessionId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ActivateSession", flags, ch, sessionId)
}

func (v *seat) ActivateSession(flags dbus.Flags, sessionId string) error {
	return (<-v.GoActivateSession(flags, make(chan *dbus.Call, 1), sessionId).Done).Err
}

// method SwitchTo

func (v *seat) GoSwitchTo(flags dbus.Flags, ch chan *dbus.Call, vtnr uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchTo", flags, ch, vtnr)
}

func (v *seat) SwitchTo(flags dbus.Flags, vtnr uint32) error {
	return (<-v.GoSwitchTo(flags, make(chan *dbus.Call, 1), vtnr).Done).Err
}

// method SwitchToNext

func (v *seat) GoSwitchToNext(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchToNext", flags, ch)
}

func (v *seat) SwitchToNext(flags dbus.Flags) error {
	return (<-v.GoSwitchToNext(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SwitchToPrevious

func (v *seat) GoSwitchToPrevious(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchToPrevious", flags, ch)
}

func (v *seat) SwitchToPrevious(flags dbus.Flags) error {
	return (<-v.GoSwitchToPrevious(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property Id s

func (v *seat) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property ActiveSession (so)

func (v *seat) ActiveSession() PropSessionInfo {
	return PropSessionInfo{
		Impl: v,
		Name: "ActiveSession",
	}
}

// property CanMultiSession b

func (v *seat) CanMultiSession() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanMultiSession",
	}
}

// property CanTTY b

func (v *seat) CanTTY() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanTTY",
	}
}

// property CanGraphical b

func (v *seat) CanGraphical() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanGraphical",
	}
}

// property Sessions a(so)

func (v *seat) Sessions() PropSessionInfoSlice {
	return PropSessionInfoSlice{
		Impl: v,
		Name: "Sessions",
	}
}

// property IdleHint b

func (v *seat) IdleHint() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IdleHint",
	}
}

// property IdleSinceHint t

func (v *seat) IdleSinceHint() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IdleSinceHint",
	}
}

// property IdleSinceHintMonotonic t

func (v *seat) IdleSinceHintMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IdleSinceHintMonotonic",
	}
}

type Session struct {
	session // interface org.freedesktop.login1.Session
	proxy.Object
}

func NewSession(conn *dbus.Conn, path dbus.ObjectPath) (*Session, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Session)
	obj.Object.Init_(conn, "org.freedesktop.login1", path)
	return obj, nil
}

type session struct{}

func (v *session) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*session) GetInterfaceName_() string {
	return "org.freedesktop.login1.Session"
}

// method Terminate

func (v *session) GoTerminate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Terminate", flags, ch)
}

func (v *session) Terminate(flags dbus.Flags) error {
	return (<-v.GoTerminate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Activate

func (v *session) GoActivate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Activate", flags, ch)
}

func (v *session) Activate(flags dbus.Flags) error {
	return (<-v.GoActivate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Lock

func (v *session) GoLock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Lock", flags, ch)
}

func (v *session) Lock(flags dbus.Flags) error {
	return (<-v.GoLock(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Unlock

func (v *session) GoUnlock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unlock", flags, ch)
}

func (v *session) Unlock(flags dbus.Flags) error {
	return (<-v.GoUnlock(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetIdleHint

func (v *session) GoSetIdleHint(flags dbus.Flags, ch chan *dbus.Call, idle bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIdleHint", flags, ch, idle)
}

func (v *session) SetIdleHint(flags dbus.Flags, idle bool) error {
	return (<-v.GoSetIdleHint(flags, make(chan *dbus.Call, 1), idle).Done).Err
}

// method SetLockedHint

func (v *session) GoSetLockedHint(flags dbus.Flags, ch chan *dbus.Call, locked bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLockedHint", flags, ch, locked)
}

func (v *session) SetLockedHint(flags dbus.Flags, locked bool) error {
	return (<-v.GoSetLockedHint(flags, make(chan *dbus.Call, 1), locked).Done).Err
}

// method Kill

func (v *session) GoKill(flags dbus.Flags, ch chan *dbus.Call, who string, signo int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Kill", flags, ch, who, signo)
}

func (v *session) Kill(flags dbus.Flags, who string, signo int32) error {
	return (<-v.GoKill(flags, make(chan *dbus.Call, 1), who, signo).Done).Err
}

// method TakeControl

func (v *session) GoTakeControl(flags dbus.Flags, ch chan *dbus.Call, force bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TakeControl", flags, ch, force)
}

func (v *session) TakeControl(flags dbus.Flags, force bool) error {
	return (<-v.GoTakeControl(flags, make(chan *dbus.Call, 1), force).Done).Err
}

// method ReleaseControl

func (v *session) GoReleaseControl(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseControl", flags, ch)
}

func (v *session) ReleaseControl(flags dbus.Flags) error {
	return (<-v.GoReleaseControl(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method TakeDevice

func (v *session) GoTakeDevice(flags dbus.Flags, ch chan *dbus.Call, major uint32, minor uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TakeDevice", flags, ch, major, minor)
}

func (*session) StoreTakeDevice(call *dbus.Call) (fd dbus.UnixFD, inactive bool, err error) {
	err = call.Store(&fd, &inactive)
	return
}

func (v *session) TakeDevice(flags dbus.Flags, major uint32, minor uint32) (fd dbus.UnixFD, inactive bool, err error) {
	return v.StoreTakeDevice(
		<-v.GoTakeDevice(flags, make(chan *dbus.Call, 1), major, minor).Done)
}

// method ReleaseDevice

func (v *session) GoReleaseDevice(flags dbus.Flags, ch chan *dbus.Call, major uint32, minor uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseDevice", flags, ch, major, minor)
}

func (v *session) ReleaseDevice(flags dbus.Flags, major uint32, minor uint32) error {
	return (<-v.GoReleaseDevice(flags, make(chan *dbus.Call, 1), major, minor).Done).Err
}

// method PauseDeviceComplete

func (v *session) GoPauseDeviceComplete(flags dbus.Flags, ch chan *dbus.Call, major uint32, minor uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PauseDeviceComplete", flags, ch, major, minor)
}

func (v *session) PauseDeviceComplete(flags dbus.Flags, major uint32, minor uint32) error {
	return (<-v.GoPauseDeviceComplete(flags, make(chan *dbus.Call, 1), major, minor).Done).Err
}

// signal PauseDevice

func (v *session) ConnectPauseDevice(cb func(major uint32, minor uint32, type0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PauseDevice", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PauseDevice",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var major uint32
		var minor uint32
		var type0 string
		err := dbus.Store(sig.Body, &major, &minor, &type0)
		if err == nil {
			cb(major, minor, type0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ResumeDevice

func (v *session) ConnectResumeDevice(cb func(major uint32, minor uint32, fd dbus.UnixFD)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ResumeDevice", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ResumeDevice",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var major uint32
		var minor uint32
		var fd dbus.UnixFD
		err := dbus.Store(sig.Body, &major, &minor, &fd)
		if err == nil {
			cb(major, minor, fd)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Lock

func (v *session) ConnectLock(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Lock", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Lock",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Unlock

func (v *session) ConnectUnlock(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Unlock", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Unlock",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Id s

func (v *session) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property User (uo)

func (v *session) User() PropSessionUser {
	return PropSessionUser{
		Impl: v,
	}
}

type PropSessionUser struct {
	Impl proxy.Implementer
}

func (p PropSessionUser) Get(flags dbus.Flags) (value UserInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"User", &value)
	return
}

func (p PropSessionUser) ConnectChanged(cb func(hasValue bool, value UserInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v UserInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, UserInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"User", cb0)
}

// property Name s

func (v *session) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Timestamp t

func (v *session) Timestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Timestamp",
	}
}

// property TimestampMonotonic t

func (v *session) TimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimestampMonotonic",
	}
}

// property VTNr u

func (v *session) VTNr() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "VTNr",
	}
}

// property Seat (so)

func (v *session) Seat() PropSessionSeat {
	return PropSessionSeat{
		Impl: v,
	}
}

type PropSessionSeat struct {
	Impl proxy.Implementer
}

func (p PropSessionSeat) Get(flags dbus.Flags) (value SeatInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Seat", &value)
	return
}

func (p PropSessionSeat) ConnectChanged(cb func(hasValue bool, value SeatInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v SeatInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, SeatInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Seat", cb0)
}

// property TTY s

func (v *session) TTY() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "TTY",
	}
}

// property Display s

func (v *session) Display() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Display",
	}
}

// property Remote b

func (v *session) Remote() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Remote",
	}
}

// property RemoteHost s

func (v *session) RemoteHost() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RemoteHost",
	}
}

// property RemoteUser s

func (v *session) RemoteUser() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RemoteUser",
	}
}

// property Service s

func (v *session) Service() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Service",
	}
}

// property Desktop s

func (v *session) Desktop() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Desktop",
	}
}

// property Scope s

func (v *session) Scope() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Scope",
	}
}

// property Leader u

func (v *session) Leader() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Leader",
	}
}

// property Audit u

func (v *session) Audit() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Audit",
	}
}

// property Type s

func (v *session) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

// property Class s

func (v *session) Class() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Class",
	}
}

// property Active b

func (v *session) Active() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Active",
	}
}

// property State s

func (v *session) State() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "State",
	}
}

// property IdleHint b

func (v *session) IdleHint() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IdleHint",
	}
}

// property IdleSinceHint t

func (v *session) IdleSinceHint() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IdleSinceHint",
	}
}

// property IdleSinceHintMonotonic t

func (v *session) IdleSinceHintMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IdleSinceHintMonotonic",
	}
}

// property LockedHint b

func (v *session) LockedHint() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "LockedHint",
	}
}

type User struct {
	user // interface org.freedesktop.login1.User
	proxy.Object
}

func NewUser(conn *dbus.Conn, path dbus.ObjectPath) (*User, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(User)
	obj.Object.Init_(conn, "org.freedesktop.login1", path)
	return obj, nil
}

type user struct{}

func (v *user) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*user) GetInterfaceName_() string {
	return "org.freedesktop.login1.User"
}

// method Terminate

func (v *user) GoTerminate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Terminate", flags, ch)
}

func (v *user) Terminate(flags dbus.Flags) error {
	return (<-v.GoTerminate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Kill

func (v *user) GoKill(flags dbus.Flags, ch chan *dbus.Call, signo int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Kill", flags, ch, signo)
}

func (v *user) Kill(flags dbus.Flags, signo int32) error {
	return (<-v.GoKill(flags, make(chan *dbus.Call, 1), signo).Done).Err
}

// property UID u

func (v *user) UID() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "UID",
	}
}

// property GID u

func (v *user) GID() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "GID",
	}
}

// property Name s

func (v *user) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Timestamp t

func (v *user) Timestamp() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Timestamp",
	}
}

// property TimestampMonotonic t

func (v *user) TimestampMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimestampMonotonic",
	}
}

// property RuntimePath s

func (v *user) RuntimePath() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "RuntimePath",
	}
}

// property Service s

func (v *user) Service() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Service",
	}
}

// property Slice s

func (v *user) Slice() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Slice",
	}
}

// property Display (so)

func (v *user) Display() PropSessionInfo {
	return PropSessionInfo{
		Impl: v,
		Name: "Display",
	}
}

// property State s

func (v *user) State() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "State",
	}
}

// property Sessions a(so)

func (v *user) Sessions() PropSessionInfoSlice {
	return PropSessionInfoSlice{
		Impl: v,
		Name: "Sessions",
	}
}

// property IdleHint b

func (v *user) IdleHint() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IdleHint",
	}
}

// property IdleSinceHint t

func (v *user) IdleSinceHint() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IdleSinceHint",
	}
}

// property IdleSinceHintMonotonic t

func (v *user) IdleSinceHintMonotonic() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "IdleSinceHintMonotonic",
	}
}

// property Linger b

func (v *user) Linger() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Linger",
	}
}

type PropSessionInfo struct {
	Impl proxy.Implementer
	Name string
}

func (p PropSessionInfo) Get(flags dbus.Flags) (value SessionInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropSessionInfo) Set(flags dbus.Flags, value SessionInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropSessionInfo) ConnectChanged(cb func(hasValue bool, value SessionInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v SessionInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, SessionInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

type PropSessionInfoSlice struct {
	Impl proxy.Implementer
	Name string
}

func (p PropSessionInfoSlice) Get(flags dbus.Flags) (value []SessionInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropSessionInfoSlice) Set(flags dbus.Flags, value []SessionInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropSessionInfoSlice) ConnectChanged(cb func(hasValue bool, value []SessionInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []SessionInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}
