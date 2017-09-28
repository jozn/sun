package models

import "ms/sun/models/x"

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
	panic("implement me")
}

func (rpcSync) SetLastSyncGeneralUpdateId(i *x.PB_SyncParam_SetLastSyncGeneralUpdateId, p x.RPC_UserParam) (*x.PB_SyncResponse_SetLastSyncGeneralUpdateId, error) {
	panic("implement me")
}

func (rpcSync) SetLastSyncNotifyUpdateId(i *x.PB_SyncParam_SetLastSyncNotifyUpdateId, p x.RPC_UserParam) (*x.PB_SyncResponse_SetLastSyncNotifyUpdateId, error) {
	panic("implement me")
}
