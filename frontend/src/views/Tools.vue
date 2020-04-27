<template>
    <div>
        <a-row :gutter="16">
            <a-col :span="5">
                <a-card title="Tree">
                    <a-tree :treeData="treeData" :replaceFields="replaceFields" :loadData="onLoadData" @select="this.onSelect" />
                </a-card>
            </a-col>
            <a-col :span="19">
                <a-table :columns="columns" :dataSource="tableData" childrenColumnName="" size="middle" :loading="tableLoading" :pagination="false">
                    <template slot="title">
                        <a-row>
                            <a-col :span="2">
                                <a-button-group>
                                    <a-button type="primary" icon="home" @click="Home" :disabled="dataParent.length === 0" />
                                    <a-button type="primary" icon="up" @click="Up" :disabled="dataParent.length === 0" />
                                </a-button-group>
                            </a-col>
                            <a-col :span="16">
                                <b style="margin-left: 30px; font-size: 18px;">{{ table.path}}</b>
                            </a-col>
                            <a-col :span="6" align="right">
                                <a-button type="primary" icon="reload" @click="Scan">ReScan</a-button>
                            </a-col>
                        </a-row>
                    </template>
                    <template slot="name" slot-scope="text, record">
                        <div v-if="record.items"><a @click="dataSet(text)">{{text}}</a></div>
                        <div v-else>{{text}}</div>
                    </template>
                    <template slot="size" slot-scope="text">
                        <a-tooltip placement="right" :title="text">
                            {{text|bytesToSize}}
                        </a-tooltip>
                    </template>
                    <template slot="percent" slot-scope="text, record">
                        <a-progress :percent="Math.round(100 / table.size * record.size)" status="normal">
                            <template v-slot:format="percent">
                                {{percent}}%
                            </template>
                        </a-progress>
                    </template>
                </a-table>
                <router-view/>
            </a-col>
        </a-row>
    </div>
</template>

<script>
    import {HTTP} from '../main'

    export default {
        name: 'Tools',
        data: function() {
            return {
                columns: [
                    {
                        title: 'Name',
                        dataIndex: 'name',
                        key: 'name',
                        scopedSlots: { customRender: 'name' }
                    },
                    {
                        title: 'Path',
                        dataIndex: 'path',
                    },
                    {
                        title: 'Size',
                        dataIndex: 'size',
                        key: 'size',
                        sorter: (a, b) => a.size - b.size,
                        sortDirections: ['descend'],
                        sortOrder: 'descend',
                        scopedSlots: { customRender: 'size' }
                    },
                    {
                        title: 'Items',
                        dataIndex: 'items',
                    },
                    {
                        title: 'Percent',
                        key: 'percent',
                        scopedSlots: { customRender: 'percent' },
                        width: '15%'
                    },
                ],
                table: [],
                tableData: [],
                dataParent: [],
                tableLoading: true,
                treeData: [],
                replaceFields: {
                    title: 'name',
                    key: 'path',
                }
            }
        },
        filters: {
            bytesToSize: function (bytes) {
                var sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
                if (bytes == 0) return '0 B';
                var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)));
                return Math.round(bytes / Math.pow(1024, i), 2) + ' ' + sizes[i];
            }
        },
        mounted() {
            this.Get();
            this.DirList();
        },
        methods: {
            DirList() {
                HTTP.get('api/tools/dirtree')
                    .then(response => {
                            this.treeData = response.data
                    })
            },
            Get() {
                this.tableLoading = true
                this.table = []
                HTTP.post('api/tools/get', {path: this.dataParent})
                    .then(response => {
                        if (response.data) {
                            this.table = response.data
                            this.tableData = this.table['children']
                        }
                        this.tableLoading = false
                    })
            },
            Scan() {
                this.tableLoading = true
                this.table = []
                HTTP.get('api/tools/scan')
                    .then(response => {
                        this.table = response.data
                        this.tableData = this.table['children']
                        this.tableLoading = false
                    })
            },
            Set() {
                this.tableData = this.table[0]['children']
            },
            Home() {
                this.dataParent = []
                this.Get()
            },
            Up() {
                let arr = this.dataParent
                arr.splice(-1, 1);
                this.Get()
            },
            dataSet(key) {
                this.dataParent.push(key)
                this.Get()
            },
            onSelect(selectedKeys) {
                this.dataParent = selectedKeys[0].substr(1).split('/')
                this.Get()
            },
            onLoadData(treeNode) {
                return new Promise(resolve => {
                    if (treeNode.dataRef.children) {
                        resolve();
                        return;
                    }
                    HTTP.get('api/tools/dirtree?path='+treeNode.dataRef.path)
                        .then(response => {
                            treeNode.dataRef.children = response.data
                            this.treeData = [...this.treeData];
                            resolve();
                        })
                });
            },
        }
    }
</script>

<style scoped>

</style>
