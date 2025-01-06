<template>
    <div>
        <a-row :gutter="16">
            <a-col :span="5">
                <a-card title="Tree">
                    <a-tree :treeData="treeData" :fieldNames ="replaceFields" :loadData="onLoadData" @select="onSelect" />
                </a-card>
            </a-col>
            <a-col :span="19">
                <a-table :columns="columns" :dataSource="tableData" childrenColumnName="" size="middle" :loading="tableLoading" :pagination="false">
                    <template #title>
                        <a-row>
                            <a-col flex="none">
                                <a-button-group>
                                    <a-button type="primary" @click="Home" :disabled="dataParent.length === 0">
                                        <template #icon>
                                            <HomeOutlined />
                                        </template>
                                    </a-button>
                                    <a-button type="primary" @click="Up" :disabled="dataParent.length === 0">
                                        <template #icon>
                                            <UpOutlined />
                                        </template>
                                    </a-button>
                                </a-button-group>
                            </a-col>
                            <a-col flex="auto">
                                <b style="margin-left: 30px; font-size: 18px;">{{ table.path}}</b>
                            </a-col>
                            <a-col flex="none">
                                <a-button type="primary" @click="Scan">
                                    <template #icon>
                                        <ReloadOutlined />
                                    </template>
                                    ReScan
                                </a-button>
                            </a-col>
                        </a-row>
                    </template>
                    <template #bodyCell="{ column, record }">
                        <template v-if="column.key === 'name'">
                            <div v-if="record.items"><a @click="dataSet(record.name)">{{record.name}}</a></div>
                            <div v-else>{{record.name}}</div>
                        </template>
                        <template v-else-if="column.key === 'size'">
                            <a-tooltip placement="right" :title="record.size + ' bytes'">
                                {{ $filters.bytesToSize(record.size) }}
                            </a-tooltip>
                        </template>
                        <template v-else-if="column.key === 'percent'">
                            <a-progress :percent="Math.round(100 / table.size * record.size)" status="normal">
                                <template v-slot:format="percent">
                                    {{percent}}%
                                </template>
                            </a-progress>
                        </template>
                    </template>
                </a-table>
                <router-view/>
            </a-col>
        </a-row>
    </div>
</template>

<script setup>
    import {HTTP} from '../main'
    import {onMounted, ref} from "vue";
    import {HomeOutlined, UpOutlined, ReloadOutlined} from "@ant-design/icons-vue";

    const columns = [
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
    ]

    const table = ref([])
    const tableData = ref([])
    const dataParent = ref([])
    const tableLoading = ref(true)
    const treeData = ref([])
    const replaceFields = {
        title: 'name',
        key: 'path',
    }
    const topButtonDisabled = ref(false)

    const DirList = () => {
        HTTP.get('api/tools/dirtree')
            .then(response => {
                treeData.value = response.data
            })
    }

    const Get = () => {
        tableLoading.value = true
        table.value = []
        HTTP.post('api/tools/get', {path: dataParent.value})
            .then(response => {
                if (response.data) {
                    table.value = response.data
                    tableData.value = table.value['child']
                }
                tableLoading.value = false
            })
    }

    const Scan = () => {
        tableLoading.value = true
        table.value = []
        HTTP.get('api/tools/scan')
            .then(response => {
                table.value = response.data
                tableData.value = table.value['child']
                tableLoading.value = false
            })
    }

    const Set = () => {
        tableData.value = table.value['child']
    }

    const Home = () => {
        dataParent.value = []
        Get()
    }

    const Up = () => {
        let arr = dataParent.value
        arr.splice(-1, 1);
        Get()
    }

    const dataSet = (key) => {
        dataParent.value.push(key)
        Get()
    }

    const onSelect = (selectedKeys) => {
        console.log(selectedKeys)
        dataParent.value = selectedKeys[0].split('/')
        Get()
    }

    const onLoadData = (treeNode) => {
        return new Promise(resolve => {
            if (treeNode.dataRef.children) {
                resolve();
                return;
            }
            HTTP.get('api/tools/dirtree?path='+treeNode.dataRef.path)
                .then(response => {
                    treeNode.dataRef.children = response.data
                    treeData.value = [...treeData.value];
                    resolve();
                })
        });
    }

    onMounted(() => {
        Get()
        DirList()
    })
</script>