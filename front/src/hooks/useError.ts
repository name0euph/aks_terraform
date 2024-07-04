import axios from "axios";
import { useNavigate } from "react-router-dom";
import { CsrfToken } from "../types";
import useStore from "../store";

export const useError = () => {
    const navigate = useNavigate();
    const resetEditedTask = useStore((state) => state.resetEditedTask);
    const getCsrfToken = async () => {
        const { data } = await axios.get<CsrfToken>(
            `${process.env.REACT_APP_API_URL}/csrf`
        );
        axios.defaults.headers.common["X-CSRF-Token"] = data.csrf_token;
    }

    const switchErrorHandling = (msg: string) => {
        switch (msg) {
            case 'invalid csrf token':
                // CSRFトークンの取得とエラーメッセージの表示
                getCsrfToken();
                alert('CSRF token is invalid. Please try again.');
                break;

            case 'invalid or expired jwt':
                // アラートの表示し、zustandのStateをリセットしてログインページへの遷移
                alert('access token expired. Please login again.');
                resetEditedTask();
                navigate('/');
                break;

            case 'missing or malformed jwt':
                // アラートの表示し、zustandのStateをリセットしてログインページへの遷移
                alert('access token is not valid. Please login again.');
                resetEditedTask();
                navigate('/');
                break;

            case 'duplicated key not allowed':
                // アラートの表示
                alert('email already exist. Please use another email.');
                break;

            case 'crypto/bcrypt: hashedPassword is not the hash of the given password':
                // アラートの表示
                alert('password is not correct.');
                break;

            case 'record not found':
                // アラートの表示
                alert('email is not correct.');
                break;

            default:
                alert(msg);
        }
    }
    return { switchErrorHandling };
}