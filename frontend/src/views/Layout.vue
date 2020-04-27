<template>
    <a-locale-provider :locale="en_US">
    <div>
        <a-layout>
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
                                <template slot="title">
                                    Logout
                                </template>
                                <a-icon type="poweroff" />
                            </a-tooltip>
                            </a-menu-item>
                    </a-menu>
                </a-layout-header>
                <a-layout-content :style="{ background: '#fff', padding: '24px', minHeight: '280px' }"><router-view/></a-layout-content>
                <a-layout-footer style="text-align: center">
                    WebDU ©2020 Created by <a href="https://github.com/SiTiSem" target="_blank">SiTiSem</a>
                </a-layout-footer>
            </a-layout>
        </a-layout>
    </div>
    </a-locale-provider>
</template>

<script>
    import en_US from 'ant-design-vue/lib/locale-provider/en_US';

    export default {
        name: "Layout",
        data() {
            return {
                en_US,
            };
        },
        methods: {
            Logout() {
                localStorage.removeItem('tks')
                this.$router.push('/login')
            },
        }
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
