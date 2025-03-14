//go:build windows

package edge

import (
	"unsafe"
)

type _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type iCoreWebView2CreateCoreWebView2ControllerCompletedHandler struct {
	vtbl *_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl
	impl _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl
}

func (i *iCoreWebView2CreateCoreWebView2ControllerCompletedHandler) AddRef() uint32 {
	ret, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))

	return uint32(ret)
}

func _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownQueryInterface(this *iCoreWebView2CreateCoreWebView2ControllerCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownAddRef(this *iCoreWebView2CreateCoreWebView2ControllerCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownRelease(this *iCoreWebView2CreateCoreWebView2ControllerCompletedHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInvoke(this *iCoreWebView2CreateCoreWebView2ControllerCompletedHandler, errorCode uintptr, createdController *ICoreWebView2Controller) uintptr {
	return this.impl.CreateCoreWebView2ControllerCompleted(errorCode, createdController)
}

type _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl interface {
	_IUnknownImpl
	CreateCoreWebView2ControllerCompleted(errorCode uintptr, createdController *ICoreWebView2Controller) uintptr
}

var _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerFn = _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInvoke),
}

func newICoreWebView2CreateCoreWebView2ControllerCompletedHandler(impl _ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl) *iCoreWebView2CreateCoreWebView2ControllerCompletedHandler {
	return &iCoreWebView2CreateCoreWebView2ControllerCompletedHandler{
		vtbl: &_ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerFn,
		impl: impl,
	}
}
