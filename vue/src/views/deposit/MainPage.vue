<template>
    <div class="animated fadeIn">
    <b-row>
        <b-col col="12">
            <h1 style="margin-bottom:50px; margin-top:50px; text-align:center">전체 진행상황</h1>        
        </b-col>
    </b-row>
    
    <b-row>
        <b-col col="4" style="text-align:center">
            <div class="card text-white bg-danger mb-3" style="max-width: 100%;">
                <div class="card-header">입금확인</div>
                <div class="card-body ">
                    <h4 class="card-title">미확인 입금 개수 : {{buys.Confirm}}개</h4>                    
                    <a class="btn btn-light"><router-link to="/depositlist" style="color:#f86c6b">입금확인 이동</router-link></a>
                    
                </div>
            </div>
        </b-col>
        <b-col col="4" style="text-align:center">
            <div class="card text-white bg-warning mb-3" style="max-width: 100%;">
                <div class="card-header">입금승인</div>
                <div class="card-body">
                    <h4 class="card-title">미승인 입금 개수 : {{buys.Acknowledgment}}개</h4>
                    <a class="btn btn-light"><router-link to="/depositchecklist" style="color:#ffc107">입금승인 이동</router-link></a>
                </div>
            </div>
        </b-col>
        <b-col col="4" style="text-align:center">
            <div class="card text-white bg-success mb-3" style="max-width: 100%;">
                <div class="card-header">입금완료</div>
                <div class="card-body">
                    <h4 class="card-title">완료된 입금 개수 : {{buys.Complete}}개</h4>                    
                    <a class="btn btn-light" ><router-link to="/depositconfirmlist" style="color:#4dbd74">입금완료 이동</router-link></a>
                </div>
            </div>
        </b-col>
    </b-row>
    <hr>
    <b-row style="margin-top:30px">
        <b-col col="4" style="text-align:center">
            <div class="card text-white bg-danger mb-3" style="max-width: 100%;">
                <div class="card-header">출금확인</div>
                <div class="card-body">
                    <h4 class="card-title">미확인 출금 개수 : {{sells.Confirm}}개</h4>                    
                    <a class="btn btn-light"><router-link to="/withdrawlist" style="color:#f86c6b" >출금확인 이동</router-link></a>
                </div>
            </div>
        </b-col>
        <b-col col="4" style="text-align:center">
            <div class="card text-white bg-warning mb-3" style="max-width: 100%;">
                <div class="card-header">출금승인</div>
                <div class="card-body">
                    <h4 class="card-title">미승인 출금 개수 : {{sells.Acknowledgment}}개</h4>                    
                    <a class="btn btn-light"><router-link to="/withdrawchecklist" style="color:#ffc107">출금승인 이동</router-link></a>
                </div>
            </div>
        </b-col>
        <b-col col="4" style="text-align:center">
            <div class="card text-white bg-success mb-3" style="max-width: 100%;">
                <div class="card-header">출금완료</div>
                <div class="card-body">
                    <h4 class="card-title">완료된 출금 개수 : {{sells.Complete}}개</h4>                    
                    <a class="btn btn-light" ><router-link to="/withdrawconfirmlist" style="color:#4dbd74">출금완료 이동</router-link></a>
                </div>
            </div>
        </b-col>
    </b-row>
    </div>
</template>

<script>

export default {
    data(){
        return{
            buys:[],
            sells:[]
        }
    },
    mounted(){    
        //세션의 토큰 , 아이디값이 서버에 저장된 값과 일치할때만
        
        // if(sessionStorage.getItem('token') == null || sessionStorage.getItem('id') == null){
        // this.$router.push('/loginPage')
        // }
        
        // const URI = 'http://172.30.1.46:8091/loginCheck'          

        // this.$axios.post(`${URI}` , {
        // 'id' : sessionStorage.getItem('id') , 
        // 'token' : sessionStorage.getItem('token')
        // })
        // .then((result) => {
        // console.log(result)
        // if(result.data == 'WARING'){        
        //     this.$router.push('/loginPage')
        // }
        // })


        //내용 불러오기
        this.$axios.post('http://172.30.1.46:8091/mainListCount')
        .then((result) => {      
            console.log(result)      
            this.buys = result.data.Buys
            this.sells = result.data.Sells
        })
        
    },

}
</script>

<style scoped>
  .fade-enter-active {
    transition: all .3s ease;
  }
  .fade-leave-active {
    transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
  }
  .fade-enter, .fade-leave-to {
    transform: translateX(10px);
    opacity: 0;
  }
/* .card {
    border: 0px;
} */
</style>
