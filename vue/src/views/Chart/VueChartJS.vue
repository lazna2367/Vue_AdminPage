<template>
    <section class="container">        
        <div class="columns">
          <!-- 
          <div class="column">
            <h3>Line Chart</h3>
            <line-chart></line-chart>
          </div>
           -->
           
          <div class="column" style="background:white">            
            <BarChart/>
          </div>

          <div v-if="flag" style="background:white">
            <apexchart width="1000" type="bar" :options="chartOptions" :series="series" style="color:black"></apexchart>
          </div>
        </div>

        <!-- 
        <div class="columns">
          <div class="column">
            <h3>Bubble Chart</h3>            
            <BubbleChart/>
          </div>
          <div class="column">
            <h3>Reactivity - Live update upon change in datasets</h3>            
            <Reactive/>
          </div>        
        </div>
         -->
         <div>
           <h1>google</h1>
           <GChart/>
         </div>

    </section>
</template>

<script>
import LineChart from './LineChart'
import BarChart from './BarChart'
import BubbleChart from './BubbleChart'
import Reactive from './Reactive'
import GChart from './GoogleCharts'

export default {
    data(){
        return{
          Today:[],
          DayAgo:[],
          TwoDayAgo:[],
          ThreeDayAgo:[],
          FourDayAgo:[],
          FiveDayAgo:[],
          SixDayAgo:[],

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
              flag: false
              }
    },
    components: {
        LineChart,
        BarChart,
        BubbleChart,
        Reactive,
        GChart
    },
    created(){      
      this.chartOptions.xaxis.categories[6] = this.$moment().format("YYYY-MM-DD")
      this.chartOptions.xaxis.categories[5] = this.$moment().add('day',-1).format("YYYY-MM-DD")
      this.chartOptions.xaxis.categories[4] = this.$moment().add('day',-2).format("YYYY-MM-DD")
      this.chartOptions.xaxis.categories[3] = this.$moment().add('day',-3).format("YYYY-MM-DD")
      this.chartOptions.xaxis.categories[2] = this.$moment().add('day',-4).format("YYYY-MM-DD")
      this.chartOptions.xaxis.categories[1] = this.$moment().add('day',-5).format("YYYY-MM-DD")
      this.chartOptions.xaxis.categories[0] = this.$moment().add('day',-6).format("YYYY-MM-DD")      
      
      const URI = 'http://172.30.1.46:8091/mainWeeklyCount'
        this.$axios.post(URI)
        .then((result) => {            
            this.Today = result.data.Today
            this.DayAgo = result.data.DayAgo
            this.TwoDayAgo = result.data.TwoDayAgo
            this.ThreeDayAgo = result.data.ThreeDayAgo
            this.FourDayAgo = result.data.FourDayAgo
            this.FiveDayAgo = result.data.FiveDayAgo
            this.SixDayAgo = result.data.SixDayAgo                    

            this.series[0].data[0] = this.SixDayAgo.Buy
            this.series[0].data[1] = this.FiveDayAgo.Buy
            this.series[0].data[2] = this.FourDayAgo.Buy
            this.series[0].data[3] = this.ThreeDayAgo.Buy
            this.series[0].data[4] = this.TwoDayAgo.Buy
            this.series[0].data[5] = this.DayAgo.Buy
            this.series[0].data[6] = this.Today.Buy

            this.series[1].data[0] = this.SixDayAgo.Sell
            this.series[1].data[1] = this.FiveDayAgo.Sell
            this.series[1].data[2] = this.FourDayAgo.Sell
            this.series[1].data[3] = this.ThreeDayAgo.Sell
            this.series[1].data[4] = this.TwoDayAgo.Sell
            this.series[1].data[5] = this.DayAgo.Sell
            this.series[1].data[6] = this.Today.Sell
            this.flag = true            
        })
    }
}
</script>

<style scoped>
  ul {
    list-style-type: none;
    padding: 0;
  }
 
  li {
    display: inline-block;
    margin: 0 10px;
  }
 
  a {
    color: #42b983;
  }
</style>
