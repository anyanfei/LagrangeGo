package main

import (
	"testing"

	"github.com/Redmomn/LagrangeGo/client/wtlogin"
	"github.com/Redmomn/LagrangeGo/utils"
)

func TestSig(t *testing.T) {
	//appInfo := info.AppList["linux"]
	//deviceInfo := &info.DeviceInfo{
	//	Guid:          "cfcd208495d565ef66e7dff9f98764da",
	//	DeviceName:    "Lagrange-DCFCD07E",
	//	SystemKernel:  "Windows 10.0.22631",
	//	KernelVersion: "10.0.22631",
	//}
	//sig := info.NewSigInfo(8848)
	buf := utils.GetBytesFromHex("000000400000229000000000000000040000001577746c6f67696e2e7472616e735f656d7000000004000000000000000f4804a80100c80102d801356610e8a30000033d0203391f410812000000000000000000fe5a1c275f0d776ee84f29df7372dea6edd5a9c5342bdd9b17a1a0d3db86aaca572c8e6adcb9cf967d61ca9ae7a395f00b085e8f2f6ca226c72d98d01758f6a020b0b26d7fd5fcf9d18d659d4a3aade976a6a32869286b860d1a5048efc01f05d5de76eb22943a9378f2c3a1e0f711293797bb0c7b7a2002c37d5da692261e3ec48781b03fa3bb6e654389525923b1c3ea9930a042d5b00bf093b475920e4f908783ddf37e596fcf3820c12466cdb82db9d5cb38b9167070dd891e8efa22caf0b41bf2611d1a697b2b6617f54a5680efa9483de78887fbde8fb43025bf24a1945bcc870e11af0517b7f05ceb99a9f02024df3a6f251f0271766549ea0d785246da6b43598103ed13b3e7aeef0fd213a1c4d155713f5205e7087ca9129b576cd239fcc8f0c8ef685c2696c355389623053687caf26581b201a377838b06a5ead656376bc991eb04610be6c0ce205c5bbeb1a13b45534fc8202782ed87f4ba73a409f26808e1f5de4f186325ca95b98b00048f2a86883dadadf8553714ed67352b2fc42667442dd73c74fa252c841fa02e270e6a27bebe95d4594a9b95a81305c93c68c8c5abdf08431d871b7a8e5d0223b33aef80d239cbe08e08b82ecf9e21afe14ff27df42244c8d51567e4f283e8bee73c613f82bf2d82d1366ae456641e8c8dd4e9aba6d235bc2d68fbb5d7af019dca2d1709cebdd210c9064fe06611cd8863fe1b45f749a86e0eac8f188b3d9bdbc2a2e2e74ddf0554ef823dfa12186ee6acc932ac001703c7032aa28d34444b51e7f913312f985f44b30001b15fb20c64fff360fac4ec919acd901476ba11a05a19b023ce2b7d79aa819d743169f430cd51a27db62a4df5e7641a2c1f64565eaab71feca5d6ddd444a94cf171a510d82aa55a6ac7da699dd7e48db7a8f032815dae72d3c32f59522e91abbb3e14c67ec98160ef72299d94b8df351ba5ed8a0df7b8c3e78d6f99e1f6ac160ec0aa458548ff94312990e78aab970c1281725005ee7a0fe96ecc51530b1a649436c5fc1fd6fbdd9d132b4073083585f581c825aecabce8ad853b9df8e4dd608e8b23b9117eaa0afacda6582cba78ababc83b62f8d541368d2bea86bf89e878d753c3211e0965826340da9ed2fd03")
	_, err := wtlogin.ParseSSOFrame(buf, true)
	if err != nil {
		t.Errorf("err %s\n", err)
	}

}
