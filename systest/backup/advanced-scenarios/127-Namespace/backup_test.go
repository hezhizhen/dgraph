package main

import (
	"testing"

	e2eCommon "github.com/dgraph-io/dgraph/graphql/e2e/common"
	utilsCommon "github.com/dgraph-io/dgraph/systest/backup/common"
	"github.com/dgraph-io/dgraph/testutil"
	"github.com/dgraph-io/dgraph/x"
)

const (
	accessJwtHeader = "X-Dgraph-AccessToken"
	restoreLocation = "/data/backups/"
	backupDst       = "/data/backups/"
)

func Test127PlusNamespaces(t *testing.T) {
	jwtTokenAlpha1Np0, headerAlpha1Np0 := utilsCommon.GetJwtTokenAndHeader(t, "alpha1", 0)
	jwtTokenAlpha2Np0, headerAlpha2Np0 := utilsCommon.GetJwtTokenAndHeader(t, "alpha2", 0)
	_ = e2eCommon.CreateNamespaces(t, headerAlpha1Np0, "alpha1", 50)
	utilsCommon.AddItemSchema(t, headerAlpha1Np0, "alpha1")
	e2eCommon.AssertGetGQLSchema(t, testutil.ContainerAddr("alpha1", 8080), headerAlpha1Np0)
	utilsCommon.AddItem(t, 1, 50, jwtTokenAlpha1Np0, "alpha1")
	utilsCommon.CheckItemExists(t, 30, jwtTokenAlpha1Np0, "alpha1")
	utilsCommon.TakeBackup(t, jwtTokenAlpha1Np0, backupDst, "alpha1")
	utilsCommon.RunRestore(t, jwtTokenAlpha2Np0, restoreLocation, "alpha2")
	dg1 := testutil.DgClientWithLogin(t, "groot", "password", x.GalaxyNamespace)
	testutil.WaitForRestore(t, dg1, testutil.ContainerAddr("alpha2", 8080))
	e2eCommon.AssertGetGQLSchema(t, testutil.ContainerAddr("alpha2", 8080), headerAlpha2Np0)
	utilsCommon.CheckItemExists(t, 30, jwtTokenAlpha2Np0, "alpha2")
	_ = e2eCommon.CreateNamespaces(t, headerAlpha1Np0, "alpha1", 50)
	jwtTokenAlpha1Np51, headerAlpha1Np51 := utilsCommon.GetJwtTokenAndHeader(t, "alpha1", 0)
	utilsCommon.AddItemSchema(t, headerAlpha1Np51, "alpha1")
	e2eCommon.AssertGetGQLSchema(t, testutil.ContainerAddr("alpha1", 8080), headerAlpha1Np51)
	utilsCommon.AddItem(t, 51, 100, jwtTokenAlpha1Np51, "alpha1")
	utilsCommon.CheckItemExists(t, 70, jwtTokenAlpha1Np51, "alpha1")
	utilsCommon.TakeBackup(t, jwtTokenAlpha1Np0, backupDst, "alpha1")
	utilsCommon.RunRestore(t, jwtTokenAlpha2Np0, restoreLocation, "alpha2")
	dg2 := testutil.DgClientWithLogin(t, "groot", "password", x.GalaxyNamespace)
	testutil.WaitForRestore(t, dg2, testutil.ContainerAddr("alpha2", 8080))
	e2eCommon.AssertGetGQLSchema(t, testutil.ContainerAddr("alpha2", 8080), headerAlpha2Np0)
	utilsCommon.CheckItemExists(t, 30, jwtTokenAlpha2Np0, "alpha2")
	jwtTokenAlpha2Np51, headerAlpha2Np51 := utilsCommon.GetJwtTokenAndHeader(t, "alpha1", 0)
	e2eCommon.AssertGetGQLSchema(t, testutil.ContainerAddr("alpha2", 8080), headerAlpha2Np51)
	utilsCommon.CheckItemExists(t, 70, jwtTokenAlpha2Np51, "alpha2")
	_ = e2eCommon.CreateNamespaces(t, headerAlpha1Np0, "alpha1", 30)
	jwtTokenAlpha1Np130, headerAlpha1Np130 := utilsCommon.GetJwtTokenAndHeader(t, "alpha1", 0)
	utilsCommon.AddItemSchema(t, headerAlpha1Np130, "alpha1")
	e2eCommon.AssertGetGQLSchema(t, testutil.ContainerAddr("alpha1", 8080), headerAlpha1Np130)
	utilsCommon.AddItem(t, 101, 130, jwtTokenAlpha1Np130, "alpha1")
	utilsCommon.CheckItemExists(t, 110, jwtTokenAlpha1Np130, "alpha1")
	utilsCommon.TakeBackup(t, jwtTokenAlpha1Np0, backupDst, "alpha1")
	utilsCommon.RunRestore(t, jwtTokenAlpha2Np0, restoreLocation, "alpha2")
	dg3 := testutil.DgClientWithLogin(t, "groot", "password", x.GalaxyNamespace)
	testutil.WaitForRestore(t, dg3, testutil.ContainerAddr("alpha2", 8080))
	e2eCommon.AssertGetGQLSchema(t, testutil.ContainerAddr("alpha2", 8080), headerAlpha2Np0)
	utilsCommon.CheckItemExists(t, 30, jwtTokenAlpha2Np0, "alpha2")
	e2eCommon.AssertGetGQLSchema(t, testutil.ContainerAddr("alpha2", 8080), headerAlpha2Np51)
	utilsCommon.CheckItemExists(t, 70, jwtTokenAlpha2Np51, "alpha2")
	jwtTokenAlpha2Np130, headerAlpha2Np130 := utilsCommon.GetJwtTokenAndHeader(t, "alpha1", 0)
	e2eCommon.AssertGetGQLSchema(t, testutil.ContainerAddr("alpha2", 8080), headerAlpha2Np130)
	utilsCommon.CheckItemExists(t, 110, jwtTokenAlpha2Np130, "alpha2")
}
