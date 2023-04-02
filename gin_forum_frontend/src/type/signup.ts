export interface ISignupData{
    username:string,
    password:string,
    re_password:string,
}

export class SignupData{
    ruleForm: ISignupData = {
        username: "",
        password: "",
        re_password:""
    }
}