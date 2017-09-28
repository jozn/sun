package models

import (
	"ms/sun/base"
	"ms/sun/models/x"
)

type rpcSync int

func (rpcSync) GetDirectUpdates(i *x.PB_SyncParam_GetDirectUpdates, p x.RPC_UserParam) (*x.PB_SyncResponse_GetDirectUpdates, error) {
	return DirectLog_GetSync(p.GetUserId(), int(i.LastId))
}

func (rpcSync) GetGeneralUpdates(i *x.PB_SyncParam_GetGeneralUpdates, p x.RPC_UserParam) (*x.PB_SyncResponse_GetGeneralUpdates, error) {
	return GeneralLog_GetLastView(p.GetUserId(), int(i.LastId))
}

func (rpcSync) GetNotifyUpdates(i *x.PB_SyncParam_GetNotifyUpdates, p x.RPC_UserParam) (*x.PB_SyncResponse_GetNotifyUpdates, error) {
	panic("implement me")
}

func (rpcSync) SetLastSyncDirectUpdateId(i *x.PB_SyncParam_SetLastSyncDirectUpdateId, p x.RPC_UserParam) (*x.PB_SyncResponse_SetLastSyncDirectUpdateId, error) {
	x.NewSession_Updater().LastSyncDirectUpdateId(int(i.LastId)).UserId_Eq(p.GetUserId()).Update(base.DB)
	return &x.PB_SyncResponse_SetLastSyncDirectUpdateId{}, nil
}

func (rpcSync) SetLastSyncGeneralUpdateId(i *x.PB_SyncParam_SetLastSyncGeneralUpdateId, p x.RPC_UserParam) (*x.PB_SyncResponse_SetLastSyncGeneralUpdateId, error) {
	x.NewSession_Updater().LastSyncGeneralUpdateId(int(i.LastId)).UserId_Eq(p.GetUserId()).Update(base.DB)
	return &x.PB_SyncResponse_SetLastSyncGeneralUpdateId{}, nil
}

func (rpcSync) SetLastSyncNotifyUpdateId(i *x.PB_SyncParam_SetLastSyncNotifyUpdateId, p x.RPC_UserParam) (*x.PB_SyncResponse_SetLastSyncNotifyUpdateId, error) {
	x.NewSession_Updater().LastSyncNotifyUpdateId(int(i.LastId)).UserId_Eq(p.GetUserId()).Update(base.DB)
	return &x.PB_SyncResponse_SetLastSyncNotifyUpdateId{}, nil
}
