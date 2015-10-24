package widgets

//#include "qgesturerecognizer.h"
import "C"
import (
	"unsafe"
)

type QGestureRecognizer struct {
	ptr unsafe.Pointer
}

type QGestureRecognizerITF interface {
	QGestureRecognizerPTR() *QGestureRecognizer
}

func (p *QGestureRecognizer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGestureRecognizer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGestureRecognizer(ptr QGestureRecognizerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGestureRecognizerPTR().Pointer()
	}
	return nil
}

func QGestureRecognizerFromPointer(ptr unsafe.Pointer) *QGestureRecognizer {
	var n = new(QGestureRecognizer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGestureRecognizer) QGestureRecognizerPTR() *QGestureRecognizer {
	return ptr
}

//QGestureRecognizer::ResultFlag
type QGestureRecognizer__ResultFlag int

var (
	QGestureRecognizer__Ignore           = QGestureRecognizer__ResultFlag(0x0001)
	QGestureRecognizer__MayBeGesture     = QGestureRecognizer__ResultFlag(0x0002)
	QGestureRecognizer__TriggerGesture   = QGestureRecognizer__ResultFlag(0x0004)
	QGestureRecognizer__FinishGesture    = QGestureRecognizer__ResultFlag(0x0008)
	QGestureRecognizer__CancelGesture    = QGestureRecognizer__ResultFlag(0x0010)
	QGestureRecognizer__ResultState_Mask = QGestureRecognizer__ResultFlag(0x00ff)
	QGestureRecognizer__ConsumeEventHint = QGestureRecognizer__ResultFlag(0x0100)
	QGestureRecognizer__ResultHint_Mask  = QGestureRecognizer__ResultFlag(0xff00)
)