<template>
    <div>
        <!-- <div>
            계정 출력
            <div v-for="value in accounts" v-bind:key="value.id">
            {{value}}
            </div>
        </div>
        <br>
        <div>
            코인베이스 계정 출력 : {{etherbase}}
        </div>
        <br>
        <div>
            코인베이스의 이더량을 출력 : {{getbal}}
        </div>
        <br>
        <div>
            현재까지 채굴 된 블록의 갯수를 표시 : {{mined}}
        </div> -->

        <div>
            발신자 주소<br>
            <input type="text" v-model="toAccounts"/>            
        </div>

        <div>
            수신자 주소<br>
            <input type="text" v-model="toAddress"/>            
        </div>

        <div>
            개수<br>
            <input type="text" v-model="sendAmountVal"/>            
        </div>
        <button v-on:click="addressBtn">전송</button>
        <br>
        <br>
        <div>
            발신자가 보낸 거래수량 : {{b_nonce}} <br>
            블록해쉬값 : {{b_blockHash}} <br>        
            블록넘버 : {{b_blockNumber}} <br>
            발신자 주소 : {{b_from}} <br>
            수신자 주소 : {{b_to}} <br>
            전송된 가격 : {{b_value / 1000000000000000000}}ETH <br>
            발신자가 제공 한 가스 : {{b_gasPrice}} <br>
        </div>
    </div>
</template>

<script>
import Web3 from 'web3'
var web3 = new Web3();
web3.setProvider(new Web3.providers.HttpProvider('http://172.30.1.27:7545'))                        

//let myContract = new web3.eth.Contract(Mycontract, '')
//필터설정
// var blockFilter = web3.eth.filter('latest')
// console.log(blockFilter)
// blockFilter.watch(function(error, blockHash) {
//     console.log(blockHash)
//     var block = web3.eth.getBlock(blockHash);    
//     console.log("block.number : " + block.number)
//     console.log("block.hash : " + block.hash)    
//     console.log("block.transactions.length : " + block.transactions.length)    
//     console.log("gasUsed : " + block.gasUsed)    
//     })    



export default {
    data(){
        return{
            accounts: [],
            etherbase: '',
            getbal: '',
            mined: '',

            toAccounts:'',
            toAddress:'',
            sendAmountVal:'',
            sendAmount:'',
            blocktext:'',

            b_nonce : '',
            b_blockHash : '',
            b_blockNumber : '',
            b_from : '',
            b_to : '',
            b_value : '',
            b_gasPrice : '',
        }
    },
    created(){
        this.web3Run()
    },
    methods:{
        web3Run(){            
            
            //모든 계정 출력
            for(var i=0; i<web3.eth.accounts.length; i++){
                this.accounts[i] = web3.eth.accounts[i]
            }

            //코인베이스 계정 출력
            this.etherbase = web3.eth.coinbase

            //코인베이스의 이더량을 출력
            this.getbal = web3.eth.getBalance(web3.eth.coinbase)

            //현재까지 채굴 된 블록의 갯수를 표시
            // web3.eth.getBlockNumber(
            //     function(a, blocknum) {
            //         this.mined = blocknum
            //     }
            // )

            //받을 코인

            //보낼사람
            
        },
        addressBtn(){
            this.sendAmount = web3.toWei(this.sendAmountVal , 'ether')
            var txHash = web3.eth.sendTransaction({
                from: this.toAccounts,
                to: this.toAddress,
                value: this.sendAmount
            })            
            
            console.log(txHash)                                    

            var hash = web3.eth.getTransaction(txHash)
            this.b_nonce = hash.nonce
            this.b_blockHash = hash.blockHash
            this.b_blockNumber = hash.blockNumber 
            this.b_from = hash.from
            this.b_to = hash.to
            this.b_value = hash.value
            this.b_gasPrice = hash.gas
            // console.log("hash : " + hash.hash)
            // console.log("nonce : " + hash.nonce)
            // console.log("blockHash : " + hash.blockHash)
            // console.log("blockNumber : " + hash.blockNumber)
            // console.log("transactionIndex : " + hash.transactionIndex)
            // console.log("from : " + hash.from)
            // console.log("to : " + hash.to)
            // console.log("value : " + hash.value)
            // console.log("gasPrice : " + hash.gasPrice)
            // console.log("gas : " + hash.gas)
            // console.log("input : " + hash.input)

        }
        
    }
    
    

}
</script>

<style>

</style>
