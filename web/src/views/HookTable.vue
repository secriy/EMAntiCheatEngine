<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item> <i class="el-icon-lx-cascades"></i> 应用日志 </el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <el-row class="handle-box">
                <el-col :span="12">
                    <el-input v-model="searchData" placeholder="玩家搜索" class="handle-input mr20" clearable></el-input
                ></el-col>
                <el-col :span="12"
                    ><el-switch
                        class="mr20"
                        style="display: block"
                        v-model="select"
                        active-color="#13ce66"
                        inactive-color="#ff4949"
                        active-text="显示全部"
                        inactive-text="显示异常"
                    >
                    </el-switch
                ></el-col>
            </el-row>
            <el-table
                :data="
                    tableData === null
                        ? tableData
                        : tableData
                              .filter(data => !searchData || data.player.toLowerCase().includes(searchData.toLowerCase()))
                              .slice((query.pageIndex - 1) * query.pageSize, query.pageIndex * query.pageSize)
                "
                border
                class="table"
                ref="multipleTable"
                header-cell-class-name="table-header"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55" align="center"></el-table-column>
                <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
                <el-table-column label="IP" width="120" :show-overflow-tooltip="true">
                    <template slot-scope="scope">{{ scope.row.ip }}</template>
                </el-table-column>
                <el-table-column label="玩家" width="140" :show-overflow-tooltip="true">
                    <template slot-scope="scope">{{ scope.row.player }}</template>
                </el-table-column>
                <el-table-column label="内容" :show-overflow-tooltip="true">
                    <template slot-scope="scope">{{ scope.row.content }}</template>
                </el-table-column>
                <el-table-column label="创建时间">
                    <template slot-scope="scope"> {{ timestampToTime(scope.row.created_at) }}</template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination
                    background
                    layout="total, prev, pager, next"
                    :current-page="query.pageIndex"
                    :page-size="query.pageSize"
                    :total="pageTotal"
                    :hide-on-single-page="hideOnSinglePage"
                    @current-change="handlePageChange"
                ></el-pagination>
            </div>
        </div>
    </div>
</template>

<script>
import * as API from '@/api/hook.js';

export default {
    data() {
        return {
            select: true,
            searchData: '',
            query: {
                pageIndex: 1,
                pageSize: 10
            },
            tableData: [],
            multipleSelection: [],
            pageTotal: 0,
            form: {},
            idx: -1,
            id: -1,
            hideOnSinglePage: true
        };
    },
    created() {
        this.select == true ? this.getData() : this.getExpData();
    },
    watch: {
        // 监听分页
        pageTotal() {
            if (this.pageTotal == (this.query.pageIndex - 1) * this.query.pageSize && this.pageTotal != 0) {
                this.query.pageIndex -= 1;
                this.select == true ? this.getData() : this.getExpData();
            }
        },
        // 监听选项
        select() {
            this.select == true ? this.getData() : this.getExpData();
        }
    },
    methods: {
        getData() {
            API.hookList().then(res => {
                this.tableData = res.data;
                this.pageTotal = res.data !== null ? res.data.length : 0;
            });
        },
        getExpData() {
            API.hookExpList().then(res => {
                this.tableData = res.data;
                this.pageTotal = res.data !== null ? res.data.length : 0;
            });
        },
        // 分页导航
        handlePageChange(val) {
            this.$set(this.query, 'pageIndex', val);
            this.select == true ? this.getData() : this.getExpData();
        },
        // 时间格式化
        timestampToTime(time, format = 'YY-MM-DD hh:mm:ss') {
            var date = new Date(time * 1000);
            var year = date.getFullYear(),
                month = date.getMonth() + 1, //月份是从0开始的
                day = date.getDate(),
                hour = date.getHours(),
                min = date.getMinutes(),
                sec = date.getSeconds();
            var preArr = Array.apply(null, Array(10)).map(function(elem, index) {
                return '0' + index;
            }); //开个长度为10的数组 格式为 ["00", "01", "02", "03", "04", "05", "06", "07", "08", "09"]

            var newTime = format
                .replace(/YY/g, year)
                .replace(/MM/g, preArr[month] || month)
                .replace(/DD/g, preArr[day] || day)
                .replace(/hh/g, preArr[hour] || hour)
                .replace(/mm/g, preArr[min] || min)
                .replace(/ss/g, preArr[sec] || sec);

            return newTime;
        }
    }
};
</script>

<style scoped>
.handle-box {
    margin-bottom: 20px;
}

.handle-input {
    width: 300px;
    display: inline-block;
}
.table {
    width: 100%;
    font-size: 14px;
}
.red {
    color: #ff0000;
}
.mr10 {
    margin-right: 10px;
}
.mr20 {
    margin-left: 50px;
}
.table-td-thumb {
    display: block;
    margin: auto;
    width: 40px;
    height: 40px;
}
</style>
<style>
.el-tooltip__popper {
    width: 50%;
}
</style>
