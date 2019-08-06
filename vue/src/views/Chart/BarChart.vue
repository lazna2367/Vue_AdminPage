<script>
//Importing Bar class from the vue-chartjs wrapper
import {Bar} from 'vue-chartjs'
//Exporting this so it can be used in other components

var date = new Date(); 
var year = date.getFullYear(); 
var month = new String(date.getMonth()+1); 
var day = new String(date.getDate()); 

if(month.length == 1){ 
  month = "0" + month; 
} 
if(day.length == 1){ 
  day = "0" + day; 
} 

var today = year + "/" + month + "/" + day


export default {    
  extends: Bar,
  data() {
    return {        

      Today:[],
      DayAgo:[],
      TwoDayAgo:[],
      ThreeDayAgo:[],
      FourDayAgo:[],
      FiveDayAgo:[],
      SixDayAgo:[],

      datacollection: {
        //Data to be represented on x-axis
        labels: ['D-6','D-5','D-4','D-3','D-2','D-1',today],
        datasets: [
          {
            label: '입금',
            backgroundColor: '#6f42c1',
            pointBackgroundColor: 'white',
            borderWidth: 1,
            pointBorderColor: '#249EBF',
            //Data to be represented on y-axis
            data: [10,20,30,40,50,60,70]
          },
          {
            label: '출금',
            backgroundColor: '#4dbd74',
            pointBackgroundColor: 'white',
            borderWidth: 1,
            pointBorderColor: '#249EBF',
            //Data to be represented on y-axis
            data: [10,20,30,40,50,60,70]
          }
        ]        
      },
      //Chart.js options that controls the appearance of the chart
      options: {
        scales: {
          yAxes: [{
            ticks: {
              beginAtZero: true
            },
            gridLines: {
              display: true
            }
          }],
          xAxes: [ {
            gridLines: {
              display: false
            }
          }]
        },
        legend: {
            display: true
          },
        responsive: true,
        maintainAspectRatio: false
      }
    }
  },

  methods: {
      // reload() {
      //       const URI = 'http://172.30.1.46:8091/mainWeeklyCount'
      //       this.$axios.post(URI)
      //       .then((result) => {
      //           console.log(result)             
      //       }
      //     )}                  
  },
  created(){
          const URI = 'http://172.30.1.46:8091/mainWeeklyCount'
          this.$axios.post(URI)
          .then((result) => {
              console.log(result)    
              this.Today = result.data.Today
              this.DayAgo = result.data.DayAgo
              this.TwoDayAgo = result.data.TwoDayAgo
              this.ThreeDayAgo = result.data.ThreeDayAgo
              this.FourDayAgo = result.data.FourDayAgo
              this.FiveDayAgo = result.data.FiveDayAgo
              this.SixDayAgo = result.data.SixDayAgo
              this.datacollection.datasets[0].data[0] = this.SixDayAgo.Buy
              this.datacollection.datasets[0].data[1] = this.FiveDayAgo.Buy
              this.datacollection.datasets[0].data[2] = this.FourDayAgo.Buy
              this.datacollection.datasets[0].data[3] = this.ThreeDayAgo.Buy
              this.datacollection.datasets[0].data[4] = this.TwoDayAgo.Buy
              this.datacollection.datasets[0].data[5] = this.DayAgo.Buy
              this.datacollection.datasets[0].data[6] = this.Today.Buy

              this.datacollection.datasets[1].data[0] = this.SixDayAgo.Sell
              this.datacollection.datasets[1].data[1] = this.FiveDayAgo.Sell
              this.datacollection.datasets[1].data[2] = this.FourDayAgo.Sell
              this.datacollection.datasets[1].data[3] = this.ThreeDayAgo.Sell
              this.datacollection.datasets[1].data[4] = this.TwoDayAgo.Sell
              this.datacollection.datasets[1].data[5] = this.DayAgo.Sell
              this.datacollection.datasets[1].data[6] = this.Today.Sell
              

              this.renderChart(this.datacollection, this.options)

          })

    // var datasetA = this.datacollection.datasets[0].data[0]
    // datasetA.set(1)
    // console.log(datasetA)
    
  }
  
  
}
</script>