<template>
  <div class="animated fadeIn">

    <b-row style="word-break: break-all; text-align: center; ">
      <b-col col-12 col-sm-12 col-md-12 col-lg-12 col-xl-12>
        <c-table :table-data="items" :fields="fields" caption="<i class='icon-wallet'></i> 입금확인 관리 테이블"></c-table>
      </b-col>

    </b-row><!--/.row-->

  </div>

</template>

<script>
import { shuffleArray } from '@/shared/utils'
import cTable from '../deposit/DepositListTable.vue'

export default {  
  components: {cTable},
  data: () => {
    return {
      
      items:[],
      // items: someData, 
      //itemsArray: someData(),
      
      fields: [
        {key: 'Id' , label: '고유값' ,sortable: true}, //고유번호,
        {key: 'Wallet.Wallet_address' , label: '지갑주소'}, //신청자 지갑주소
        {key: 'Base_unit' , label: '입금된 화폐종류'}, //입금할 코인종류
        {key: 'Target_unit' , label: '변경할 화폐종류'}, // 변경할 코인종류
        {key: 'Base_amount' , label: '입금된 화폐개수'}, // 입금 코인개수
        {key: 'Target_amount' , label: '받을 화폐개수'}, // 받을 코인개수
        {key: 'Rate_by_base' , label: '환전시세'}, //환전시세
        {key: 'Payout_account.Transaction' , label: '트랜잭션'}, //거래내역
        {key: 'Payout_account.OutcomAddress' , label: '회사 지갑주소'}, // 회사 지갑주소
        {key: 'Created_at' , label: '신청시간' , sortable: true}, // 신청시간
        {key: 'Updated_at' , label: '업데이트 시간'}, //승인시간
        {key: 'Confirm' , label: '승인여부' , sortable: true }, //승인여부        
        
        //{key: 'username', label: 'User', sortable: true},
        //{key: 'registered'},
        //{key: 'role'},
        //{key: 'status', sortable: true}
      ],
    }
  },  
  mounted(){    
    //세션의 토큰 , 아이디값이 서버에 저장된 값과 일치할때만
    
    if(sessionStorage.getItem('token') == null || sessionStorage.getItem('id') == null){
      this.$router.push('/loginPage')
    }
    
    const URI = 'http://172.30.1.46:8091/loginCheck'          

    this.$axios.post(`${URI}` , {
      'id' : sessionStorage.getItem('id') , 
      'token' : sessionStorage.getItem('token')
    })
    .then((result) => {
      console.log(result)
      if(result.data == 'WARING'){        
        this.$router.push('/loginPage')
      }
    })
    
  },
  created(){
    this.axiosload()
  },
  methods:{
    // axiosload() {
    //   console.log('도니?1');
    //   this.$store.dispatch('updateClick')
    //   console.log(this.$store.state.items)
    // }
    axiosload() {
      const baseURI = 'http://172.30.1.46:8091/depositList'
      this.$axios.post(`${baseURI}`)
      .then((result) => {
        console.log(result) 
        if(result.data != null){
          this.items = result.data      
        } else if(result.data == null){
          this.items = []
        }
        
      })
    }

  }
}
</script>
