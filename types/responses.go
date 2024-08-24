package types

type ServerCode int

const (
  ServerCodeOK = ServerCode(200)
  ServerCodeNotFound = ServerCode(404)
  ServerCodeInsufficientQuantity = ServerCode(478)
  ServerCodeActionInProgress = ServerCode(486)
  ServerCodeCharacterAlreadyAtPosition = ServerCode(490)
  ServerCodeNotSkillLevelRequired = ServerCode(493)
  ServerCodeInventoryFull = ServerCode(497)
  ServerCodePlayerNotFound = ServerCode(498)
  ServerCodeCharacterInCooldown = ServerCode(499)
  ServerCodeResourceNotFoundOnMap = ServerCode(598)
)

type ServerCodeInfo struct {
  Type ServerCode
}

type Response struct {
  ServerCodeInfo
  Data []byte
}
