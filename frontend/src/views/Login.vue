<template>
        <a-flex justify="center" align="center" style="height: 100vh">
                <a-card title="Login" style="max-width: 400px; width: 400px;">
                    <a-alert
                        v-if="errorValue"
                        message="Login or password is incorrect"
                        type="error"
                        closable
                        :afterClose="handleClose"
                    />
                    <a-form ref="formRef" :model="formState" @submit="handleSubmit" :rules="rules">
                        <a-form-item ref="username" name="username">
                            <a-input
                                v-model:value="formState.username"
                                placeholder="Username"
                            >
                                <template #prefix><user-outlined /></template>
                            </a-input>
                        </a-form-item>
                        <a-form-item ref="password" name="password">
                            <a-input
                                v-model:value="formState.password"
                                type="password"
                                placeholder="Password"
                            >
                                <template #prefix><lock-outlined /></template>
                            </a-input>
                        </a-form-item>
                        <a-form-item>
                            <a-button :disabled="disabled" type="primary" html-type="submit" block>
                                Login
                            </a-button>
                        </a-form-item>
                    </a-form>
                    <div style="text-align: center;">
                        WebDU ©{{ new Date().getFullYear() }} Created by <a href="https://github.com/SiTiSem" target="_blank">SiTiSem</a>
                    </div>
                </a-card>
        </a-flex>
</template>

<script setup>
import {useRouter} from "vue-router";
import {computed, reactive, ref} from "vue";
import {UserOutlined, LockOutlined} from '@ant-design/icons-vue';
import axios from 'axios'
import {Form} from "ant-design-vue";

const useForm = Form.useForm;

const router = useRouter()

const errorValue = ref(false);

const formRef = ref();

const formState = reactive({
    username: '',
    password: '',
});

const rules = reactive({
    username: [
        { required: true, message: 'Please input your username!', trigger: 'blur' },
    ],
    password: [
        { required: true, message: 'Please input your password!', trigger: 'blur' },
    ],
})

const disabled = computed(() => {
    return !(formState.username && formState.password);
});

const { resetFields, validate, validateInfos } = useForm(formState, rules, {
    onValidate: (...args) => console.log(...args),
});

const handleSubmit = () => {
    validate()
        .then(() => {
            axios.post('api/login', {
                            username: formState.username,
                            password: formState.password
                        })
                            .then(response => {
                                localStorage.setItem('tks', response.data.token);
                                router.push({path: '/'})
                            })
                            .catch(error => {
                                errorValue.value = true
                            })
        })
        .catch(err => {
            errorValue.value = true
        });

}

const handleClose = () => {
    errorValue.value = false
}


</script>

<style scoped>
body {
    height: 100vh;
}
.ant-alert {
    margin-bottom: 20px;
}
</style>
