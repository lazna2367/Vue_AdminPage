<template>
  <div class="app flex-row align-items-center">
    <div class="container">
      <b-row class="justify-content-center">
        <b-col md="8">
          <b-card-group>
            <b-card no-body class="p-4">
              <b-card-body>
                <b-form>
                  <h1>관리자 로그인</h1>
                  <p class="text-muted">아이디와 비밀번호를 입력해주세요.</p>
                  <b-input-group class="mb-3">
                    <b-input-group-prepend><b-input-group-text><i class="icon-user"></i></b-input-group-text></b-input-group-prepend>
                    <b-form-input type="text" class="form-control" placeholder="아이디" autocomplete="username email" v-model="user.ID" v-on:keyup.enter="submitBtn"/>
                  </b-input-group>
                  <b-input-group class="mb-4">
                    <b-input-group-prepend><b-input-group-text><i class="icon-lock"></i></b-input-group-text></b-input-group-prepend>
                    <b-form-input type="password" class="form-control" placeholder="비밀번호" autocomplete="current-password" v-model="user.PW" v-on:keyup.enter="submitBtn"/>                    
                  </b-input-group>
                  <b-alert v-show="alert.show" show variant="danger">{{alert.Msg}}</b-alert>
                  <b-row>
                    <b-col cols="12" class="text-center">
                      <b-button variant="primary" class="px-8" v-on:click="submitBtn">로그인</b-button>
                    </b-col>
                  </b-row>                  
                </b-form>
              </b-card-body>
            </b-card>            

          </b-card-group>
        </b-col>
      </b-row>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return{
      user: {
        ID : '',
        PW: '',
      },
      resultToken: '',
      resultFail: '',
      alert:{
        Msg : '',
        show : false
      }

    }
  },
  methods:{
    submitBtn(){
      const baseURI = 'http://172.30.1.46:8091/login'
      
      this.$axios.post(baseURI , this.user)
      .then((result) => {
        console.log(result)
        this.resultToken = result.data.Token
        console.log(this.resultToken)
        if(this.resultToken == undefined){
          this.resultFail = result.data
          if(this.resultFail == 'No ID'){
            this.alert.show = true
            this.alert.Msg = '아이디를 찾을 수 없습니다.'
            console.log(this.alert.show)
            //alert('아이디를 찾을 수 없습니다.')
          }else if (this.resultFail == 'Miss PW'){
            this.alert.show = true
            this.alert.Msg = '비밀번호를 확인해주세요.'
            //alert('비밀번호를 확인해주세요.')
          }
        } else {
          sessionStorage.setItem('id' , this.user.ID)
          sessionStorage.setItem('token' , this.resultToken)
          this.$router.push('/')
        }
      })

    }
  },

  mounted(){
    if(sessionStorage.getItem('id') != null && sessionStorage.getItem('token') != null){      
      //세션의 토큰 , 아이디값이 서버에 저장된 값이 존재할때
      const URI = 'http://172.30.1.46:8091/loginCheck'          
      this.$axios.post(`${URI}` , {
        'id' : sessionStorage.getItem('id') , 
        'token' : sessionStorage.getItem('token')
      })
      .then((result) => {
        if(result.data == 'SUCCESS'){        
          this.$router.push('/')
        }
      })
    }


  }


}
</script>
<style>
/* ../../../public/img/01.png */
.align-items-center{
  background-image: url('../../../public/img/02.jpg');
  background-repeat: no-repeat;
  background-position: center center;
  background-size: 100%;
  
}

</style>

