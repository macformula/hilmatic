package speedgoat

import (
	"go.uber.org/zap"
	"net"
)

const (
	_digitalPinCount   = 16
	_analogOutputCount = 4
	_analogInputCount  = 8
	_analogPinCount    = 12
	_loggerName        = "speedgoat_controller"
)

// Controller provides control for various Speedgoat pins
type Controller struct {
	ip   net.IP
	port int
	l    *zap.Logger

	digital [_digitalPinCount]bool
	analog  [_analogPinCount]float64
}

// NewController returns a new Speedgoat controller
func NewController(l *zap.Logger) *Controller {
	sg := Controller{
		l: l,
	}
	return &sg
}

// Open configures the controller
func (c *Controller) Open() error {
	return nil
}

// SetDigital sets an output digital pin for a Speedgoat digital pin
func (c *Controller) SetDigital(output *DigitalPin, b bool) error {
	return nil
}

// ReadDigital returns the level of a Speedgoat digital pin
func (c *Controller) ReadDigital(output *DigitalPin) (bool, error) {
	return false, nil
}

// WriteVoltage sets the voltage of a Speedgoat analog pin
func (c *Controller) WriteVoltage(output *AnalogPin, voltage float64) error {
	return nil
}

// ReadVoltage returns the voltage of a Speedgoat analog pin
func (c *Controller) ReadVoltage(output *AnalogPin) (float64, error) {
	return 0.00, nil
}

// WriteCurrent sets the current of a Speedgoat analog pin
func (c *Controller) WriteCurrent(output *AnalogPin, current float64) error {
	return nil
}

// ReadCurrent returns the current of a Speedgoat analog pin
func (c *Controller) ReadCurrent(output *AnalogPin) (float64, error) {
	return 0.00, nil
}

func (c *Controller) tickOutputs() {
	// call a pack function for the digital and analog arrays here, transmit every 10 milliseconds
}

func (c *Controller) tickInputs() {
	// call unpack here on digital and analog arrays, receive every 10 milliseconds
	// if we have not received a tcp packet in over a second, error out
}
