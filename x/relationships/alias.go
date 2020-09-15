package relationships

// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/relationships/client/cli"
	"github.com/desmos-labs/desmos/x/relationships/client/rest"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types/models"
	"github.com/desmos-labs/desmos/x/relationships/types/msgs"
)

const (
	ModuleName               = models.ModuleName
	RouterKey                = models.RouterKey
	StoreKey                 = models.StoreKey
	ActionCreateRelationship = models.ActionCreateRelationship
	ActionDeleteRelationship = models.ActionDeleteRelationship
	ActionBlockUser          = models.ActionBlockUser
	ActionUnblockUser        = models.ActionUnblockUser
	QuerierRoute             = models.QuerierRoute
	QueryUserRelationships   = models.QueryUserRelationships
	QueryRelationships       = models.QueryRelationships
	QueryUserBlocks          = models.QueryUserBlocks
)

var (
	// functions aliases
	GetQueryCmd                  = cli.GetQueryCmd
	GetCmdQueryRelationships     = cli.GetCmdQueryRelationships
	GetCmdQueryUserRelationships = cli.GetCmdQueryUserRelationships
	GetCmdQueryUserBlocks        = cli.GetCmdQueryUserBlocks
	GetTxCmd                     = cli.GetTxCmd
	GetCmdCreateRelationship     = cli.GetCmdCreateRelationship
	GetCmdDeleteRelationship     = cli.GetCmdDeleteRelationship
	GetCmdBlockUser              = cli.GetCmdBlockUser
	GetCmdUnblockUser            = cli.GetCmdUnblockUser
	RegisterRoutes               = rest.RegisterRoutes
	NewHandler                   = keeper.NewHandler
	NewKeeper                    = keeper.NewKeeper
	NewQuerier                   = keeper.NewQuerier
	NewRelationship              = models.NewRelationship
	NewUserBlock                 = models.NewUserBlock
	RelationshipsStoreKey        = models.RelationshipsStoreKey
	UsersBlocksStoreKey          = models.UsersBlocksStoreKey
	RegisterModelsCodec          = models.RegisterModelsCodec
	NewMsgCreateRelationship     = msgs.NewMsgCreateRelationship
	NewMsgDeleteRelationship     = msgs.NewMsgDeleteRelationship
	NewMsgBlockUser              = msgs.NewMsgBlockUser
	NewMsgUnblockUser            = msgs.NewMsgUnblockUser
	RegisterMessagesCodec        = msgs.RegisterMessagesCodec

	// variable aliases
	MsgsCodec                = msgs.MsgsCodec
	RelationshipsStorePrefix = models.RelationshipsStorePrefix
	UsersBlocksStorePrefix   = models.UsersBlocksStorePrefix
	ModelsCdc                = models.ModelsCdc
)

type (
	CommonRelationshipReq = rest.CommonRelationshipReq
	UserBlockReq          = rest.UserBlockReq
	UserUnblockReq        = rest.UserUnblockReq
	Keeper                = keeper.Keeper
	Relationship          = models.Relationship
	Relationships         = models.Relationships
	UserBlock             = models.UserBlock
	MsgCreateRelationship = msgs.MsgCreateRelationship
	MsgDeleteRelationship = msgs.MsgDeleteRelationship
	MsgBlockUser          = msgs.MsgBlockUser
	MsgUnblockUser        = msgs.MsgUnblockUser
)
