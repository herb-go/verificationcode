package verificationcode

type Status int

const StatusSuccess = Status(0)
const StatusFail = Status(-1)
const StatusUnavailableUser = Status(1)
const StatusRetryLater = Status(2)
