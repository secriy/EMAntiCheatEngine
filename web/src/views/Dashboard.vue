<template>
    <div>
        <el-row :gutter="20">
            <el-row :gutter="29" class="mgb20">
                <el-col :span="8">
                    <el-card shadow="hover" :body-style="{ padding: '0px' }">
                        <div class="grid-content grid-con-1">
                            <i class="el-icon-setting grid-con-icon"></i>
                            <div class="grid-cont-right">
                                <div class="grid-num">{{ username }}</div>
                                <div>{{ '管理员' }}</div>
                            </div>
                        </div>
                    </el-card>
                </el-col>
                <el-col :span="8" v-if="username === 'admin'">
                    <el-card shadow="hover" :body-style="{ padding: '0px' }">
                        <div class="grid-content grid-con-2">
                            <i class="el-icon-lx-people grid-con-icon"></i>
                            <div class="grid-cont-right">
                                <div class="grid-num">{{ user_num }}</div>
                                <div>总用户数</div>
                            </div>
                        </div>
                    </el-card>
                </el-col>
            </el-row>
        </el-row>
        <div>
            <div id="pieHook" style="width:400px;height:300px;float:left"></div>
            <div id="pieAnti" style="width:400px;height:300px;float:left"></div>
        </div>
    </div>
</template>

<script>
// 引入基本模板
let echarts = require('echarts/lib/echarts');
// 引入饼图组件
require('echarts/lib/chart/pie');
require('echarts/theme/macarons'); // echarts theme
import bus from '@/components/bus';
import * as U_API from '@/api/user.js';
import * as H_API from '@/api/hook.js';
import * as A_API from '@/api/anti.js';

export default {
    name: 'dashboard',
    data() {
        return {
            username: sessionStorage.getItem('username'),
            user_num: null,
            charts: '',
            option: ['正常行为', '异常行为'],
            optionData: {
                hook: [
                    { value: 0, name: '相对正常行为', itemStyle: '#1ab394' },
                    { value: 0, name: '相对异常行为', itemStyle: '#79d2c0' }
                ],
                anti: [
                    { value: 0, name: '正常行为', itemStyle: '#1ab394' },
                    { value: 0, name: '异常行为', itemStyle: '#79d2c0' }
                ]
            }
        };
    },
    mounted() {
        this.$nextTick(function(data) {
            H_API.hookCount().then(res => {
                this.drawPie('pieHook', res.data.count_all, res.data.count_exp);
            });
            A_API.antiCount().then(res => {
                this.drawPie('pieAnti', res.data.count_all, res.data.count_exp);
            });
        });
    },
    created() {
        // 用户列表
        U_API.userList().then(res => {
            this.user_num = res.data.length;
        });
    },
    activated() {
        // 用户列表
        U_API.userList().then(res => {
            this.user_num = res.data.length;
        });
    },
    methods: {
        drawPie(id, all, exp) {
            this.charts = echarts.init(document.getElementById(id), 'macarons');
            if (id == 'pieHook') {
                this.optionData.hook[0].value = all;
                this.optionData.hook[1].value = exp;
            } else {
                this.optionData.anti[0].value = all;
                this.optionData.anti[1].value = exp;
            }
            this.setOptions(id);
        },
        setOptions(id) {
            this.charts.setOption({
                title: {
                    text: id == 'pieHook' ? '应用日志' : 'Anti日志'
                },
                tooltip: {
                    trigger: 'item',
                    formatter: '{a}<br/>{b}:{c} ({d}%)'
                },
                legend: {
                    bottom: 10,
                    left: 'center',
                    data: this.opinion
                },
                series: [
                    {
                        name: '数量',
                        type: 'pie',
                        radius: '65%',
                        center: ['50%', '50%'],
                        avoidLabelOverlap: false,
                        itemStyle: {
                            emphasis: {
                                shadowBlur: 10,
                                shadowOffsetX: 0,
                                shadowColor: 'rgba(0, 0, 0, 0.5)'
                            },
                            color: function(params) {
                                //自定义颜色
                                var colorList = ['#1ab394', '#79d2c0'];
                                return colorList[params.dataIndex];
                            }
                        },
                        data: id == 'pieHook' ? this.optionData.hook : this.optionData.anti
                    }
                ]
            });
        }
    }
};
</script>

<style scoped>
.el-row {
    margin-bottom: 20px;
}

.grid-content {
    display: flex;
    align-items: center;
    height: 100px;
}

.grid-cont-right {
    flex: 1;
    text-align: center;
    font-size: 14px;
    color: #999;
}

.grid-num {
    font-size: 30px;
    font-weight: bold;
}

.grid-con-icon {
    font-size: 50px;
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    color: #fff;
}

.grid-con-1 .grid-con-icon {
    background: rgb(45, 140, 240);
}

.grid-con-1 .grid-num {
    color: rgb(45, 140, 240);
}

.grid-con-2 .grid-con-icon {
    background: rgb(100, 213, 114);
}

.grid-con-2 .grid-num {
    color: rgb(45, 140, 240);
}

.user-info {
    display: flex;
    align-items: center;
    padding-bottom: 20px;
    border-bottom: 2px solid #ccc;
    margin-bottom: 20px;
}

.user-avator {
    width: 120px;
    height: 120px;
    border-radius: 50%;
}

.user-info-cont {
    padding-left: 50px;
    flex: 1;
    font-size: 14px;
    color: #999;
}

.user-info-cont div:first-child {
    font-size: 30px;
    color: #222;
}

.user-info-list {
    font-size: 14px;
    color: #999;
    line-height: 25px;
}

.user-info-list span {
    margin-left: 70px;
}

.mgb20 {
    margin-bottom: 20px;
}

.todo-item {
    font-size: 14px;
}

.todo-item-del {
    text-decoration: line-through;
    color: #999;
}
</style>
