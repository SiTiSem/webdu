<template>
            <a-layout>
                <a-layout-header>
                    <div class="logo">WebDU</div>
                    <a-menu
                            theme="dark"
                            mode="horizontal"
                            :defaultSelectedKeys="['1']"
                            :style="{ lineHeight: '64px' }"
                    >
                        <a-menu-item @click="Logout">
                            <a-tooltip>
                                <template #title>
                                    Logout
                                </template>
                                <PoweroffOutlined />
                            </a-tooltip>
                            </a-menu-item>
                    </a-menu>
                </a-layout-header>
                <a-layout-content :style="{ background: '#fff', padding: '24px', minHeight: '280px' }"><router-view/></a-layout-content>
                <a-layout-footer style="text-align: center">
                    WebDU ©2024 Created by <a href="https://github.com/SiTiSem" target="_blank">SiTiSem</a>
                </a-layout-footer>
            </a-layout>
</template>

<script setup>
import {PoweroffOutlined} from "@ant-design/icons-vue";
import router from "@/router/index.js";

const Logout = () => {
    localStorage.removeItem('tks')
    router.push('/login')
}
</script>

<style scoped>
    .logo {
        color: white;
        font-size: 24px;
        font-weight: 600;
        margin-right: 10px;
        float: left;
    }
</style>
