<template>
    <div class="animated fadeIn">
    <div class="card border-secondary mb-3" >
        <b-row class="back_row" style="margin-left:0px;margin-right:0px;">
            <b-col xl="12" style="padding-left:30px; padding-right: 30px; padding-top:50px; padding-bottom:30px;">
                <div class="card border-secondary mb-3" style="background:#e7ecef">
                <div class="card-header bg-transparent border-secondary" style="text-align:center; margin-top: 10px;"><h1 class="hn"><i class="icon-wallet"></i> 입출금 현황판</h1></div>        
                <!-- <div class="card-header bg-transparent border-secondary" style="margin-bottom:50px; margin-top:20px; text-align:center">진행상황</div> -->
                    <b-row style="padding-left: 20px; padding-right: 20px; padding-top:40px ;padding-bottom:30px;">
                    <b-col xl="3" style="text-align:center">
                        <div class="card text-white bg-danger mb-3" style="max-width: 100%;">
                            <div class="card-header"><i class="fa fa-level-down fa-lg"/> 입금확인</div>
                            <div class="card-body" style="background:white">
                                <h4 class="card-title" style="color:#f75d5c">미확인 입금 개수 : {{buys.Confirm}}개</h4>                    
                                <a class="btn btn-danger"><router-link to="/depositlist" style="color:white">입금확인 이동</router-link></a>                    
                            </div>
                        </div>                
                        <div class="card text-white bg-warning mb-3" style="max-width: 100%;">
                            <div class="card-header"><i class="fa fa-level-down fa-lg"/> 입금승인</div>
                            <div class="card-body" style="background:white">
                                <h4 class="card-title" style="color:#f7b900">미승인 입금 개수 : {{buys.Acknowledgment}}개</h4>
                                <a class="btn btn-warning"><router-link to="/depositchecklist" style="color:white">입금승인 이동</router-link></a>
                            </div>
                        </div>
                        <div class="card text-white bg-success mb-3" style="max-width: 100%;">
                            <div class="card-header"><i class="fa fa-level-down fa-lg"/> 입금완료</div>
                            <div class="card-body" style="background:white">
                                <h4 class="card-title" style="color:#44b76c">완료된 입금 개수 : {{buys.Complete}}개</h4>                    
                                <a class="btn btn-success" ><router-link to="/depositconfirmlist" style="color:white">입금완료 이동</router-link></a>
                            </div>
                        </div>
                    </b-col>
                    
                    <b-col xl="3" style="text-align:center">
                        <div class="card text-white bg-danger mb-3" style="max-width: 100%;">
                            <div class="card-header"><i class="fa fa-level-up fa-lg"/>출금확인</div>
                            <div class="card-body" style="background:white">
                                <h4 class="card-title" style="color:#f75d5c">미확인 출금 개수 : {{sells.Confirm}}개</h4>                    
                                <a class="btn btn-danger"><router-link to="/withdrawlist" style="color:white" >출금확인 이동</router-link></a>
                            </div>
                        </div>
                        <div class="card text-white bg-warning mb-3" style="max-width: 100%;">
                            <div class="card-header"><i class="fa fa-level-up fa-lg"/>출금승인</div>
                            <div class="card-body" style="background:white">
                                <h4 class="card-title" style="color:#f7b900">미승인 출금 개수 : {{sells.Acknowledgment}}개</h4>                    
                                <a class="btn btn-warning"><router-link to="/withdrawchecklist" style="color:white">출금승인 이동</router-link></a>
                            </div>
                        </div>
                        <div class="card text-white bg-success mb-3" style="max-width: 100%;">
                            <div class="card-header"><i class="fa fa-level-up fa-lg"/>출금완료</div>
                            <div class="card-body" style="background:white">
                                <h4 class="card-title" style="color:#44b76c">완료된 출금 개수 : {{sells.Complete}}개</h4>                    
                                <a class="btn btn-success" ><router-link to="/withdrawconfirmlist" style="color:white">출금완료 이동</router-link></a>
                            </div>
                        </div>
                    </b-col>
                    <b-col xl="1">                        
                    </b-col>
                    <b-col xl="5" style="text-align:center">
                        <div class="card text-Light bg-Light mb-3" style="max-width: 100%; max-height: 100%;">
                            <div class="card-header" style="background:#20a8d8; color:white;"><i class="fa fa-sign-in fa-lg"/>들어온 코인정보</div>
                            <div class="card-body">
                                <b-row>
                                    <b-col xl="12">
                                        <div class="card bg-light mb-3" >
                                            신청수량 : {{mainInputCoin.Count}}
                                        </div>
                                    </b-col>
                                </b-row>
                                <b-row>
                                    <b-col xl="3">
                                    <div class="card bg-light mb-3" style="max-width: auto;">
                                        <div class="card-body">
                                            <h4 class="card-title">BTC</h4>
                                            <p class="card-text">{{ addComma(mainInputCoin.BTC) }} BTC</p>
                                        </div>
                                    </div>
                                    </b-col>
                                    <b-col xl="3">
                                    <div class="card bg-light mb-3" style="max-width: auto;">
                                        <div class="card-body">
                                            <h4 class="card-title">ETH</h4>
                                            <p class="card-text">{{ addComma(mainInputCoin.ETH) }} ETH</p>
                                        </div>
                                    </div>
                                    </b-col>
                                    <b-col xl="3">
                                    <div class="card bg-light mb-3" style="max-width: auto;">
                                        <div class="card-body">
                                            <h4 class="card-title">WON</h4>
                                            <p class="card-text">{{ addComma(mainInputCoin.WON) }} WON</p>
                                        </div>
                                    </div>
                                    </b-col>
                                    <b-col xl="3">
                                    <div class="card bg-light mb-3" style="max-width: auto;">
                                        <div class="card-body">
                                            <h4 class="card-title">MO</h4>
                                            <p class="card-text">{{ addComma(mainInputCoin.MO) }} MO</p>
                                        </div>
                                    </div>
                                    </b-col>
                                </b-row>
                            </div>
                        </div>
                        <div class="card text-Light bg-Light mb-3" style="max-width: 100%; max-height: 100%;">
                            <div class="card-header" style="background:#45668e; color:white"><i class="fa fa-sign-out fa-lg"/>나간 코인정보</div>
                            <div class="card-body">
                                <b-row>
                                    <b-col xl="12">
                                        <div class="card bg-light mb-3">
                                            신청수량 : {{mainOutputCoin.Count}}
                                        </div>
                                    </b-col>
                                </b-row>
                                <b-row>
                                    <b-col xl="3">
                                    <div class="card bg-light mb-3" style="max-width: auto;">
                                        <div class="card-body">
                                            <h4 class="card-title">BTC</h4>
                                            <p class="card-text">{{ addComma(mainOutputCoin.BTC) }} BTC</p>
                                        </div>
                                    </div>
                                    </b-col>
                                    <b-col xl="3">
                                    <div class="card bg-light mb-3" style="max-width: auto;">
                                        <div class="card-body">
                                            <h4 class="card-title">ETH</h4>
                                            <p class="card-text">{{ addComma(mainOutputCoin.ETH) }} ETH</p>
                                        </div>
                                    </div>
                                    </b-col>
                                    <b-col xl="3">
                                    <div class="card bg-light mb-3" style="max-width: auto;">
                                        <div class="card-body">
                                            <h4 class="card-title">WON</h4>
                                            <p class="card-text">{{ addComma(mainOutputCoin.WON) }} WON</p>
                                        </div>
                                    </div>
                                    </b-col>
                                    <b-col xl="3">
                                    <div class="card bg-light mb-3" style="max-width: auto;">
                                        <div class="card-body">
                                            <h4 class="card-title">MO</h4>
                                            <p class="card-text">{{ addComma(mainOutputCoin.MO) }} MO</p>
                                        </div>
                                    </div>
                                    </b-col>
                                </b-row>
                            </div>
                        </div>
                    </b-col>
                    </b-row>                                
                </div>        
            </b-col>             
            <!-- 출금 -->
            <!-- <b-col xl="6" style="padding-left:30px; padding-right: 30px; padding-top:50px; padding-bottom:30px;">           
                <div class="card border-secondary mb-3" style="background:#e7ecef">
                <div class="card-header bg-transparent border-secondary" style="text-align:center; margin-top: 10px;"><h1 class="hn"><i class="icon-wallet"></i> 입출금 현황판</h1></div>        
                <b-row style="padding-left: 20px; padding-right: 20px; padding-top:40px ;padding-bottom:30px;">
                </b-row>
                </div>
            </b-col> -->
        </b-row> 
        <hr>
        <!-- 그래프 -->
        <b-row>            
            <b-col xl="12" style="padding-left:30px; padding-right: 30px; padding-top:50px; padding-bottom:30px;">
                    <div class="card bg-light mb-3">
                        <b-row >
                        <b-col xl='6'>
                        <div class="card-header bg-transparent border-secondary" style="margin-bottom:50px; margin-top:20px; text-align:center"><h1 class="hn"><i class="fa fa-bar-chart fa-lg mt-4"></i> 입/출금 요청 차트</h1></div>                                                
                            <div class="card border-Primary mb-3">
                                <div v-if="flag_count" style="background:white">
                                    <apexchart type="bar" :options="chartOptions" :series="series" ></apexchart>
                                </div>
                            </div>                                         
                        </b-col>
                        <b-col xl='6'>
                        <div class="card-header bg-transparent border-secondary" style="margin-bottom:50px; margin-top:20px; text-align:center"><h1 class="hn"><i class="fa fa-bar-chart fa-lg mt-4"></i> 코인 개수 차트</h1></div>
                            <div class="card border-Primary mb-3">
                                <div v-if="flag_amount" style="background:white">
                                    <apexchart  type="bar" :options="chartOptions" :series="series" style="color:black"></apexchart>
                                </div>
                            </div>                                         
                        </b-col>                   
                        </b-row>   
                    </div>
           
            </b-col>        
        </b-row>
    </div>    
    </div>
</template>

<script>
export default {
    data(){
        return{
            buys:[],
            sells:[],

            C_Today:[],
            C_DayAgo:[],
            C_TwoDayAgo:[],
            C_ThreeDayAgo:[],
            C_FourDayAgo:[],
            C_FiveDayAgo:[],
            C_SixDayAgo:[],

            A_Today:[],
            A_DayAgo:[],
            A_TwoDayAgo:[],
            A_ThreeDayAgo:[],
            A_FourDayAgo:[],
            A_FiveDayAgo:[],
            A_SixDayAgo:[],

            mainInputCoin:[],
            mainOutputCoin:[],


            chartOptions: {
                chart: {
                    id: 'vuechart-example'
                },
                xaxis: {
                    categories: ['1','2','3','4','5','6','7']
                }
            },
            series: [{
                name: '입금',
                data: [0, 0, 0, 0, 0, 0, 0]
                }, 
                {
                  name: '출금',
                  data: [0, 0, 0, 0, 0, 0, 0]
                }],

            flag_count: false,
            flag_amount: false
            
        }
    },
    created(){
        this.chartOptions.xaxis.categories[6] = this.$moment().format("MM/DD")
        this.chartOptions.xaxis.categories[5] = this.$moment().add('day',-1).format("MM/DD")
        this.chartOptions.xaxis.categories[4] = this.$moment().add('day',-2).format("MM/DD")
        this.chartOptions.xaxis.categories[3] = this.$moment().add('day',-3).format("MM/DD")
        this.chartOptions.xaxis.categories[2] = this.$moment().add('day',-4).format("MM/DD")
        this.chartOptions.xaxis.categories[1] = this.$moment().add('day',-5).format("MM/DD")
        this.chartOptions.xaxis.categories[0] = this.$moment().add('day',-6).format("MM/DD")      
        
        this.$axios.post('http://172.30.1.46:8091/mainWeeklyCount')
        .then((result) => {            
            this.C_Today = result.data.Today
            this.C_DayAgo = result.data.DayAgo
            this.C_TwoDayAgo = result.data.TwoDayAgo
            this.C_ThreeDayAgo = result.data.ThreeDayAgo
            this.C_FourDayAgo = result.data.FourDayAgo
            this.C_FiveDayAgo = result.data.FiveDayAgo
            this.C_SixDayAgo = result.data.SixDayAgo                    

            this.series[0].data[0] = this.C_SixDayAgo.Buy
            this.series[0].data[1] = this.C_FiveDayAgo.Buy
            this.series[0].data[2] = this.C_FourDayAgo.Buy
            this.series[0].data[3] = this.C_ThreeDayAgo.Buy
            this.series[0].data[4] = this.C_TwoDayAgo.Buy
            this.series[0].data[5] = this.C_DayAgo.Buy
            this.series[0].data[6] = this.C_Today.Buy

            this.series[1].data[0] = this.C_SixDayAgo.Sell
            this.series[1].data[1] = this.C_FiveDayAgo.Sell
            this.series[1].data[2] = this.C_FourDayAgo.Sell
            this.series[1].data[3] = this.C_ThreeDayAgo.Sell
            this.series[1].data[4] = this.C_TwoDayAgo.Sell
            this.series[1].data[5] = this.C_DayAgo.Sell
            this.series[1].data[6] = this.C_Today.Sell
            this.flag_count = true            
        })

        this.$axios.post('http://172.30.1.46:8091/mainWeeklyAmount')
        .then((result) => {        
            this.A_Today = result.data.Today
            this.A_DayAgo = result.data.DayAgo
            this.A_TwoDayAgo = result.data.TwoDayAgo
            this.A_ThreeDayAgo = result.data.ThreeDayAgo
            this.A_FourDayAgo = result.data.FourDayAgo
            this.A_FiveDayAgo = result.data.FiveDayAgo
            this.A_SixDayAgo = result.data.SixDayAgo                    

            this.series[0].data[0] = this.A_SixDayAgo.Buy
            this.series[0].data[1] = this.A_FiveDayAgo.Buy
            this.series[0].data[2] = this.A_FourDayAgo.Buy
            this.series[0].data[3] = this.A_ThreeDayAgo.Buy
            this.series[0].data[4] = this.A_TwoDayAgo.Buy
            this.series[0].data[5] = this.A_DayAgo.Buy
            this.series[0].data[6] = this.A_Today.Buy

            this.series[1].data[0] = this.A_SixDayAgo.Sell
            this.series[1].data[1] = this.A_FiveDayAgo.Sell
            this.series[1].data[2] = this.A_FourDayAgo.Sell
            this.series[1].data[3] = this.A_ThreeDayAgo.Sell
            this.series[1].data[4] = this.A_TwoDayAgo.Sell
            this.series[1].data[5] = this.A_DayAgo.Sell
            this.series[1].data[6] = this.A_Today.Sell
            this.flag_amount = true            
        })

        // mainInputCoin:[],
        // mainOutputCoin:[],
        this.$axios.post('http://172.30.1.46:8091/mainInputCoin')
        .then((result) =>{
            this.mainInputCoin = result.data
        })
        this.$axios.post('http://172.30.1.46:8091/mainOutputCoin')
        .then((result) =>{
            this.mainOutputCoin = result.data
        })
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
        if(result.data == 'WARING'){        
            this.$router.push('/loginPage')
        }
        })


        //내용 불러오기
        this.$axios.post('http://172.30.1.46:8091/mainListCount')
        .then((result) => {                  
            this.buys = result.data.Buys
            this.sells = result.data.Sells
        })                
        
    },
    methods:{
        addComma(num){
            var regexp = /\B(?=(\d{3})+(?!\d))/g;
            return num.toString().replace(regexp, ',');
        }
    }

}
</script>

<style scoped>
@import url(//fonts.googleapis.com/earlyaccess/hanna.css);
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
    .row{
     margin-left: 0px;
     margin-right: 0px;   
    }
    .hn{
        font-family: 'Hanna' ,sans-serif;;
    }
</style>
