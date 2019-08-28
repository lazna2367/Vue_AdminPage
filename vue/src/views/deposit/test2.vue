<template>
    <div>
        체인정보요청 <br>
        Last_mci : {{getinfo.Last_mci}} <br>
        Last_stable_mci : {{getinfo.Last_stable_mci}} <br>
        <hr>
        새로운 주소 발급 : {{getnewaddress.Result}} <br>
        <button v-on:click="newaddress">주소발급</button> <br>
        <hr>
        주소검증 : {{validateaddress.Result}} <br>
        <input type="text" placeholder="주소입력" v-model="addresschack"><br>
        <button v-on:click="dateaddress">주소검증 확인</button> <br>
        <hr>
        잔액조회 <br>
        Stable : {{getbalance.Stable / 1000000}} <br>
        Pending : {{getbalance.Pending}} <br>
        <button v-on:click="btngetbalance">잔액조회 새로고침</button>
        <hr>
        거래내역 조회 : {{listtransactions}} <br>
        <button v-on:click="btnlisttransactions">거래내역 조회</button><br>
        <hr>
        송금 : <br>
        <input type="text" placeholder="지갑주소" v-model="address"><br>
        <input type="text" placeholder="입금량" v-model="inputText"><button v-on:click="btnsendtoaddress">송금버튼</button>
    </div>
</template>

<script>
export default {
    data(){
        return{
            getinfo:'', //체인 정보요청
            getnewaddress:'', //새로운 주소 발급
            validateaddress:'', //주소 검증
            getbalance:'', //잔액조회
            listtransactions:'', //거래내역 조회
            sendtoaddress:'', //송금
            inputText:'',
            address:'' ,
            addresschack:'', 
        }

    },
    created(){
        const URI = 'http://172.30.1.46:8091/getinfo'

        // this.$axios.post(`${URI}` , {
        //     "jsonrpc":"2.0",
        //     "id":1, 
        //     "method":"getinfo",
        //     "params":{}
        // })
        // .then((result) => {
        //     this.getinfo = result.data
        // })
        

        //getinfo
        this.$axios.post('http://172.30.1.46:8091/getinfo')
        .then((result) =>{
            this.getinfo = result.data.Result
        })

        //getbalance
        this.$axios.post('http://172.30.1.46:8091/getbalance')
        .then((result) =>{
            this.getbalance = result.data.Result.Base
        })

    },
    methods:{
        newaddress(){
            this.$axios.post('http://172.30.1.46:8091/getnewaddress')
            .then((result) =>{
                this.getnewaddress = result.data
            })        
        },
        dateaddress(){
            var newaddress = this.addresschack //this.getnewaddress.Result            
            this.$axios.post('http://172.30.1.46:8091/validateaddress' , 
            {
                "Address" : newaddress
            }
            )
            .then((result) =>{
                this.validateaddress = result.data
            }) 
        },
        btnlisttransactions(){
            var mci = this.getinfo.Result.Last_stable_mci
            
            this.$axios.post('http://172.30.1.46:8091/listtransactions',
            {
                "Mci" : mci
            }
            )
            .then((result) =>{
                this.listtransactions = result.data
            })  
        },
        btnsendtoaddress(){
            var newaddress = this.address //this.getnewaddress.Result
            var money = this.inputText * 1000000 //100만 
            this.$axios.post('http://172.30.1.46:8091/sendtoaddress' ,
            {
                "Address" : newaddress , 
                "Amount" : money
            }                
            )
            .then((result) => {
                console.log(result)
                
            })
        },
        btngetbalance(){            
            this.$axios.post('http://172.30.1.46:8091/getbalance')
            .then((result) =>{
                this.getbalance = result.data.Result.Base
            })
        }


    }
    
}
</script>

<style>

</style>
