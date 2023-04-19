package listimages

import "fmt"

type NotRunningError struct{
    OperationName string
}
func (err NotRunningError) Error() string {
    return fmt.Sprintf("%v is not running", err.OperationName)
}
