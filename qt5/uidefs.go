// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qt5

import (
	"syscall"
)

var (
	EINVAL error = syscall.EINVAL
)

// STS added EchoMode, 2013-09-18
type EchoMode int

const (
	Normal EchoMode = iota
	NoEcho
	Password
	PasswordEchoOnEdit
)

type Alignment int

const (
	AlignLeft            Alignment = 0x0001
	AlignRight           Alignment = 0x0002
	AlignHCenter         Alignment = 0x0004
	AlignJustify         Alignment = 0x0008
	AlignAbsolute        Alignment = 0x0010
	AlignHorizontal_Mask Alignment = AlignLeft | AlignRight | AlignHCenter | AlignJustify | AlignAbsolute

	AlignTop           Alignment = 0x0020
	AlignBottom        Alignment = 0x0040
	AlignVCenter       Alignment = 0x0080
	AlignVertical_Mask Alignment = AlignTop | AlignBottom | AlignVCenter

	AlignCenter Alignment = AlignVCenter | AlignHCenter
)

type Orientation int

const (
	Horizontal Orientation = 0x1
	Vertical   Orientation = 0x2
)

type TickPosition int

const (
	NoTicks TickPosition = iota
	TicksAbove
	TicksBelow
	TicksBothSides
)

type DockWidgetArea int

const (
	NoDockWidgetArea     DockWidgetArea = 0
	LeftDockWidgetArea   DockWidgetArea = 0x01
	RightDockWidgetArea  DockWidgetArea = 0x02
	TopDockWidgetArea    DockWidgetArea = 0x04
	BottomDockWidgetArea DockWidgetArea = 0x08
	AllDockWidgetAreas   DockWidgetArea = LeftDockWidgetArea | RightDockWidgetArea | TopDockWidgetArea | BottomDockWidgetArea
)

type ItemFlag int
type ItemFlags int

const (
	NoItemFlags         ItemFlag = 0
	ItemIsSelectable    ItemFlag = 1
	ItemIsEditable      ItemFlag = 2
	ItemIsDragEnabled   ItemFlag = 4
	ItemIsDropEnabled   ItemFlag = 8
	ItemIsUserCheckable ItemFlag = 16
	ItemIsEnabled       ItemFlag = 32
	ItemIsTristate      ItemFlag = 64
)

type CheckState int

const (
	Unchecked        CheckState = 0
	PartiallyChecked CheckState = 1
	Checked          CheckState = 2
)

const (
	PlainText = 0
	RichText  = 1
	AutoText  = 2
	LogText   = 3
)

const (
	LightWeight    = 25
	NormalWeight   = 50
	DemiBoldWeight = 63
	BoldWeight     = 75
	BlackWeight    = 87
)

const (
	NoModifier          = 0x00000000
	ShiftModifier       = 0x02000000
	ControlModifier     = 0x04000000
	AltModifier         = 0x08000000
	MetaModifier        = 0x10000000
	KeypadModifier      = 0x20000000
	GroupSwitchModifier = 0x40000000
	// Do not extend the mask to include 0x01000000
	KeyboardModifierMask = 0xfe000000
)

//shorter names for shortcuts
const (
	META          = MetaModifier
	SHIFT         = ShiftModifier
	CTRL          = ControlModifier
	ALT           = AltModifier
	MODIFIER_MASK = KeyboardModifierMask
	UNICODE_ACCEL = 0x00000000
)

const (
	NoButton        = 0x00000000
	LeftButton      = 0x00000001
	RightButton     = 0x00000002
	MidButton       = 0x00000004 // ### Qt 5: remove me
	MiddleButton    = MidButton
	XButton1        = 0x00000008
	XButton2        = 0x00000010
	MouseButtonMask = 0x000000ff
)

//enum MessageIcon { NoIcon, Information, Warning, Critical };
type MessageIconType int

const (
	NoIcon MessageIconType = iota
	Information
	Warning
	Critical
)

type MessageBoxIcon MessageIconType

type StandardButton int
type StandardButtons int

const (
	OkButton              StandardButton = 0x00000400
	OpenButton            StandardButton = 0x00002000
	SaveButton            StandardButton = 0x00000800
	CancelButton          StandardButton = 0x00400000
	CloseButton           StandardButton = 0x00200000
	DiscardButton         StandardButton = 0x00800000
	ApplyButton           StandardButton = 0x02000000
	ResetButton           StandardButton = 0x04000000
	RestoreDefaultsButton StandardButton = 0x08000000
	HelpButton            StandardButton = 0x01000000
	SaveAllButton         StandardButton = 0x00001000
	YesButton             StandardButton = 0x00004000
	YesToAllButton        StandardButton = 0x00008000
	NoMessageButton       StandardButton = 0x00010000
	NoToAllButton         StandardButton = 0x00020000
	AbortButton           StandardButton = 0x00040000
	RetryButton           StandardButton = 0x00080000
	IgnoreButton          StandardButton = 0x00100000
	InvalidButton         StandardButton = 0x00000000
)

type ToolButtonStyle int

const (
	ToolButtonIconOnly ToolButtonStyle = iota
	ToolButtonTextOnly
	ToolButtonTextBesideIcon
	ToolButtonTextUnderIcon
	ToolButtonFollowStyle
)

type ToolButtonPopupMode int

const (
	ToolButtonDelayedPopup ToolButtonPopupMode = iota
	ToolButtonMenuButtonPopup
	ToolButtonInstantPopup
)

type PenStyle int

const (
	NoPen PenStyle = iota
	SolidLine
	DashLine
	DotLine
	DashDotLine
	DashDotDotLine
	CustomDashLine
)

type BrushStyle int

const (
	NoBrush BrushStyle = iota
	SolidPattern
	Dense1Pattern
	Dense2Pattern
	Dense3Pattern
	Dense4Pattern
	Dense5Pattern
	Dense6Pattern
	Dense7Pattern
	HorPattern
	VerPattern
	CrossPattern
	BDiagPattern
	FDiagPattern
	DiagCrossPattern
	LinearGradientPattern
	RadialGradientPattern
	ConicalGradientPattern
	TexturePattern BrushStyle = 24
)

type TransformationMode int

const (
	FastTransformation TransformationMode = iota
	SmoothTransformation
)

type AspectRatioMode int

const (
	IgnoreAspectRatio AspectRatioMode = iota
	KeepAspectRatio
	KeepAspectRatioByExpanding
)

type SizePolicyPolicyFlag int

const (
	GrowFlag   SizePolicyPolicyFlag = 1
	ExpandFlag SizePolicyPolicyFlag = 2
	ShrinkFlag SizePolicyPolicyFlag = 4
	IgnoreFlag SizePolicyPolicyFlag = 8
)

type SizePolicyPolicy int

const (
	Fixed            SizePolicyPolicy = 0
	Minimum                           = SizePolicyPolicy(GrowFlag)
	Maximum                           = SizePolicyPolicy(ShrinkFlag)
	Preferred                         = SizePolicyPolicy(GrowFlag | ShrinkFlag)
	Expanding                         = SizePolicyPolicy(GrowFlag | ShrinkFlag | ExpandFlag)
	MinimumExpanding                  = SizePolicyPolicy(GrowFlag | ExpandFlag)
	Ignored                           = SizePolicyPolicy(ShrinkFlag | GrowFlag | IgnoreFlag)
)

type SizePolicyControlType int

const (
	ControlTypeDefaultType SizePolicyControlType = 0x00000001
	ControlTypeButtonBox   SizePolicyControlType = 0x00000002
	ControlTypeCheckBox    SizePolicyControlType = 0x00000004
	ControlTypeComboBox    SizePolicyControlType = 0x00000008
	ControlTypeFrame       SizePolicyControlType = 0x00000010
	ControlTypeGroupBox    SizePolicyControlType = 0x00000020
	ControlTypeLabel       SizePolicyControlType = 0x00000040
	ControlTypeLine        SizePolicyControlType = 0x00000080
	ControlTypeLineEdit    SizePolicyControlType = 0x00000100
	ControlTypePushButton  SizePolicyControlType = 0x00000200
	ControlTypeRadioButton SizePolicyControlType = 0x00000400
	ControlTypeSlider      SizePolicyControlType = 0x00000800
	ControlTypeSpinBox     SizePolicyControlType = 0x00001000
	ControlTypeTabWidget   SizePolicyControlType = 0x00002000
	ControlTypeToolButton  SizePolicyControlType = 0x00004000
)

type HeaderResizeMode int

const (
	InteractiveHeader      HeaderResizeMode = 0
	StretchHeader          HeaderResizeMode = 1
	FixedHeader            HeaderResizeMode = 2
	ResizeHeaderToContents HeaderResizeMode = 3
)

type IODeviceOpenMode int

const (
	OpenModeNotOpen    IODeviceOpenMode = 0x0000
	OpenModeReadOnly   IODeviceOpenMode = 0x0001
	OpenModeWriteOnly  IODeviceOpenMode = 0x0002
	OpenModeReadWrite                   = IODeviceOpenMode(OpenModeReadOnly | OpenModeWriteOnly)
	OpenModeAppend     IODeviceOpenMode = 0x0004
	OpenModeTruncate   IODeviceOpenMode = 0x0008
	OpenModeText       IODeviceOpenMode = 0x0010
	OpenModeUnbuffered IODeviceOpenMode = 0x0020
)

type FileDeviceFileError int

const (
	FileDeviceNoError          FileDeviceFileError = 0
	FileDeviceReadError        FileDeviceFileError = 1
	FileDeviceWriteError       FileDeviceFileError = 2
	FileDeviceFatalError       FileDeviceFileError = 3
	FileDeviceResourceError    FileDeviceFileError = 4
	FileDeviceOpenError        FileDeviceFileError = 5
	FileDeviceAbortError       FileDeviceFileError = 6
	FileDeviceTimeOutError     FileDeviceFileError = 7
	FileDeviceUnspecifiedError FileDeviceFileError = 8
	FileDeviceRemoveError      FileDeviceFileError = 9
	FileDeviceRenameError      FileDeviceFileError = 10
	FileDevicePositionError    FileDeviceFileError = 11
	FileDeviceResizeError      FileDeviceFileError = 12
	FileDevicePermissionsError FileDeviceFileError = 13
	FileDeviceCopyError        FileDeviceFileError = 14
)

type FileDevicePermission int
type FileDevicePermissions int

const (
	FileDeviceReadOwner  FileDevicePermission = 0x4000
	FileDeviceWriteOwner FileDevicePermission = 0x2000
	FileDeviceExeOwner   FileDevicePermission = 0x1000
	FileDeviceReadUser   FileDevicePermission = 0x0400
	FileDeviceWriteUser  FileDevicePermission = 0x0200
	FileDeviceExeUser    FileDevicePermission = 0x0100
	FileDeviceReadGroup  FileDevicePermission = 0x0040
	FileDeviceWriteGroup FileDevicePermission = 0x0020
	FileDeviceExeGroup   FileDevicePermission = 0x0010
	FileDeviceReadOther  FileDevicePermission = 0x0004
	FileDeviceWriteOther FileDevicePermission = 0x0002
	FileDeviceExeOther   FileDevicePermission = 0x0001
)

type MediaAvailabilityStatus int

const (
	MediaAvailable      MediaAvailabilityStatus = 0
	MediaServiceMissing MediaAvailabilityStatus = 1
	MediaBusy           MediaAvailabilityStatus = 2
	MediaResourceError  MediaAvailabilityStatus = 3
)

type MediaEncodingMode int

const (
	MediaConstantQualityEncoding MediaEncodingMode = 0
	MediaConstantBitRateEncoding MediaEncodingMode = 1
	MediaAverageBitRateEncoding  MediaEncodingMode = 2
	MediaTwoPassEncoding         MediaEncodingMode = 3
)

type MediaEncodingQuality int

const (
	MediaVeryLowQualityEncoding  MediaEncodingQuality = 0
	MediaLowQualityEncoding      MediaEncodingQuality = 1
	MediaNormalQualityEncoding   MediaEncodingQuality = 2
	MediaHighQualityEncoding     MediaEncodingQuality = 3
	MediaVeryHighQualityEncoding MediaEncodingQuality = 4
)

type MediaPlaylistError int

const (
	MediaPlaylistNoError              MediaPlaylistError = 0
	MediaPlaylistFormatError          MediaPlaylistError = 1
	MediaPlaylistForNotSupportedError MediaPlaylistError = 2
	MediaPlaylistNetworkError         MediaPlaylistError = 3
	MediaPlaylistAccessDeniedError    MediaPlaylistError = 4
)

type MediaPlaylistPlaybackMode int

const (
	MediaPlaylistPlayCurrentItemOnce   MediaPlaylistPlaybackMode = 0
	MediaPlaylistPlayCurrentIteminLoop MediaPlaylistPlaybackMode = 1
	MediaPlaylistPlaySequential        MediaPlaylistPlaybackMode = 2
	MediaPlaylistPlayLoop              MediaPlaylistPlaybackMode = 3
	MediaPlaylistPlayRandom            MediaPlaylistPlaybackMode = 4
)

type Point struct {
	X, Y int
}

func (p Point) Unpack() (int, int) {
	return p.X, p.Y
}

type PointF struct {
	X, Y float64
}

func (p PointF) Unpack() (float64, float64) {
	return p.X, p.Y
}

type Size struct {
	Width, Height int
}

func (p Size) Unpack() (int, int) {
	return p.Width, p.Height
}

type SizeF struct {
	Width, Height float64
}

func (p SizeF) Unpack() (float64, float64) {
	return p.Width, p.Height
}

type Rect struct {
	X, Y          int
	Width, Height int
}

func (p Rect) Unpack() (int, int, int, int) {
	return p.X, p.Y, p.Width, p.Height
}

func (p Rect) Point() Point {
	return Point{p.X, p.Y}
}

func (p Rect) Size() Size {
	return Size{p.Width, p.Height}
}

type RectF struct {
	X, Y          float64
	Width, Height float64
}

func (p RectF) Unpack() (float64, float64, float64, float64) {
	return p.X, p.Y, p.Width, p.Height
}

func (p RectF) PointF() PointF {
	return PointF{p.X, p.Y}
}

func (p RectF) SizeF() SizeF {
	return SizeF{p.Width, p.Height}
}

type Margins struct {
	Left, Top, Right, Bottom int
}

func (p Margins) Unpack() (int, int, int, int) {
	return p.Left, p.Top, p.Right, p.Bottom
}

func Pt(x, y int) Point {
	return Point{x, y}
}

func PtF(x, y float64) PointF {
	return PointF{x, y}
}

func Sz(w, h int) Size {
	return Size{w, h}
}

func SzF(w, h float64) SizeF {
	return SizeF{w, h}
}

func Rc(x, y, w, h int) Rect {
	return Rect{x, y, w, h}
}

func RcF(x, y, w, h float64) RectF {
	return RectF{x, y, w, h}
}

func Mr(left, right, top, bottom int) Margins {
	return Margins{left, right, top, bottom}
}
