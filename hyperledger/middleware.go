package hyperledger


var fabric Fabric

func WriteTrans(key string, value string) string {
	rwset1, rwset2 := fabric.WriteTranaction(key, value, fabric.MSP_org1)
	if rwset1.msp == fabric.MSP_peer1 && rwset2.msp == fabric.MSP_peer2 {
		msps := [] string {rwset1.msp , rwset2.msp}
		rwset := RWSet{ key:key, value: value, peers_msp: msps}
		fabric.SendToOrderer(rwset)
		return "ok"
	}

	return "failed"
}

func GetTrans(key string) string{
	rwset := fabric.ReadTranaction(key, fabric.MSP_org1)
	return rwset
}

func StartFabric(){
	fabric = Fabric{}
	fabric.Start()
}
