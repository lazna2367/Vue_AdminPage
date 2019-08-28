<template>
    <div>         
        체인정보요청  : {{getinfo.result.last_stable_mci}} <br>
        새로운 주소 발급  : {{getnewaddress}} <br>
        <button v-on:click="newaddress">주소발급</button><br>
        주소검증  : {{validateaddress.result}} <br>
        잔액조회  : {{}} <br>
        거래내역 조회 : {{listtransactions}}<br>
        <button v-on:click="sendto">거래내역</button>
    </div>
</template>

<script>
export default {
    data(){
        return{
            getinfo:[],
            getnewaddress:'',
            validateaddress:[],
            getbalance:[],
            test:[],
            listtransactions:[],
        
        }
    },
    created(){
        //method
        //체인 정보 요청 getinfo  / params x
        //새로운 주소 발급 getnewaddress / params x
        //주소 검증 validateaddress / params x
        //잔액조회 getbalance / params :["주소"]
        //거래내역 조회 listtransactions / params  {since_mci : 1234}
        //송금 sendtoaddress / params : ["주소" , 3000000]

        const URI = 'http://testnet.dubu4.com:6552'

        this.$axios.post(`${URI}` , {
            "jsonrpc":"2.0",
            "id":1, 
            "method":"getinfo",
            "params":{}
        })
        .then((result) => {
            this.getinfo = result.data
        })

    },
    methods:{
        newaddress(){
        const URI = 'http://testnet.dubu4.com:6552'
        var newadd = ''
        //주소생성
        this.$axios.post(`${URI}` , {
            "jsonrpc":"2.0",
            "id":1, 
            "method":"getnewaddress",
            "params":{}
        })
        .then((result) => {
            this.getnewaddress = result.data.result            
            newadd = result.data.result
        })

        this.test = 
            {
            "jsonrpc":"2.0",
            "id":1, 
            "method":"validateaddress",
            "params":[`${this.getnewaddress}`]
            }
        console.log(this.getnewaddress)
        console.log(this.test)
        console.log("확인" + newadd)
        //주소검증
        this.$axios.post(`${URI}` , {
            "jsonrpc":"2.0",
            "id":1, 
            "method":"validateaddress",
            "params":[`"${this.getnewaddress.result}"`]
        })
        .then((result) => {
            this.validateaddress = result.data
        })


        },        
        sendto(){
            const URI = 'http://testnet.dubu4.com:6552'               
            var list = this.getinfo.result.last_stable_mci            
            console.log(list)   
            //주소생성
            this.$axios.post(`${URI}` , {
                "jsonrpc":"2.0",
                "id":1, 
                "method":"listtransactions",
                "params":{"since_mci" : `"${list}"`}
            })
            .then((result) => {
                this.listtransactions = result.data                            
                console.log(this.listtransactions)
            })
        }
    }   
}
</script>

<style>

</style>
