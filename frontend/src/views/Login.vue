<template>
    <div style="height: 100%; display: flex; justify-content: space-around;">
        <a-row type="flex" justify="space-around" align="middle">
            <a-col>
                <a-card title="Login" style="max-width: 400px; width: 400px;">
                    <a-alert
                            v-if="errorValue"
                            message="Login or password is incorrect"
                            type="error"
                            closable
                            :afterClose="handleClose"
                    />
                    <a-form :form="form" @submit="handleSubmit">
                        <a-form-item>
                            <a-input
                                    v-decorator="['username', { rules: [{ required: true, message: 'Please input your username!' }] },]"
                                    placeholder="Username"
                            >
                                <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)" />
                            </a-input>
                        </a-form-item>
                        <a-form-item>
                            <a-input
                                    v-decorator="['password', { rules: [{ required: true, message: 'Please input your Password!' }] },]"
                                    type="password"
                                    placeholder="Password"
                            >
                                <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)" />
                            </a-input>
                        </a-form-item>
                        <a-form-item align="center">
                            <a-button type="primary" html-type="submit" block>
                                Login
                            </a-button>
                        </a-form-item>
                    </a-form>
                    <div style="text-align: center;">
                        WebDU ©2020 Created by <a href="https://github.com/SiTiSem" target="_blank">SiTiSem</a>
                    </div>
                </a-card>
            </a-col>
        </a-row>
    </div>

</template>

<script>
    import axios from 'axios'
    export default {
        name: "Login",
        data: function() {
            return {
                errorValue: false,
            }
        },
        beforeCreate() {
            this.form = this.$form.createForm(this, { name: 'normal_login' });
        },
        methods: {
            handleSubmit(e) {
                e.preventDefault();
                this.form.validateFields((err, value) => {
                    if (!err) {
                        axios.post('/api/login', {
                            username: value.username,
                            password: value.password
                        })
                            .then(response => {
                                // Cookies.set('tks', response.data, { expires: 1 })
                                localStorage.setItem('tks', response.data.token);
                                this.$router.push({path: '/'})
                            })
                            .catch(error => {
                                this.error = error
                                this.errorValue = true
                            })
                    }
                })
            },
            handleClose() {
                this.errorValue = false;
            },
        }
    }
</script>

<style scoped>
    .ant-alert {
        margin-bottom: 20px;
    }
</style>
