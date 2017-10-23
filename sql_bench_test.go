package main

import (
	"sync"
)

var onc sync.Once

func prepareTest() {
	onc.Do(func() { startApp() })
}

/*

func Benchmark_sql_Activity_selext(b *testing.B) {
	prepareTest()

	ids := MemoryStore.UserFollowingList_Get(5).Elements
	for i := 0; i < b.N; i++ {
		NewActivity_Selector().ActorUserId_In(ids).OrderBy_Id_Desc().GetRows(base.DB)
	}
}

func Benchmark_sql_Activity_insert(b *testing.B) {
	prepareTest()

	//ids:= MemoryStore.UserFollowingList_Get(5).Elements
	for i := 0; i < b.N; i++ {
		arr := make([]Activity, 0, 1000)
		for i := 0; i < 1000; i++ {
			arr = append(arr, Activity{
				ActorUserId:  25,
				ActionTypeId: 592,
				TargetId:     rand.Intn(10000000),
				RefId:        rand.Intn(10000000),
				CreatedAt:    helper.TimeNow(),
			})
		}
		MassInsert_Activity(arr, base.DB)
	}
}

func Benchmark_sql_Insert(b *testing.B) {
	prepareTest()

	for i := 0; i < b.N; i++ {
		f := FollowingListMemberHistory{}
		f.Save(base.DB)
	}
}

func Benchmark__sql_UserById(b *testing.B) {
	prepareTest()
	for i := 0; i < b.N; i++ {
		UserById(base.DB, i)
	}
}

func Benchmark__sql_NewUser_Selector_GetRows(b *testing.B) {
	prepareTest()
	for i := 0; i < b.N; i++ {
		NewUser_Selector().GetRows2(base.DB)
	}
}
*/
