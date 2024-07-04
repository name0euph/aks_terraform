import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import useStore from "../store";
import { Credential } from "../types";
import { useError } from "../hooks/useError";

export const useMutateAuth = () => {
    const navigate = useNavigate();
    const resetEditedTask = useStore((state) => state.resetEditedTask);
    const { switchErrorHandling } = useError();

    const loginMutation = useMutation(
        async (user: Credential) =>
            await axios.post(`${process.env.REACT_APP_API_URL}/login`, user),
        {
            // ログイン成功時には、/todoに遷移
            onSuccess: () => {
            navigate('/todo');
            },
            onError: (err: any) => {
                // CSRFミドルウェアのエラーはerror.response.data.messageに格納される
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message);
                } else {
                    switchErrorHandling(err.response.data);
                }
            },
        }
    );
    const registerMutation = useMutation(
        async (user: Credential) =>
            await axios.post(`${process.env.REACT_APP_API_URL}/signup`, user),
        {
            onError: (err: any) => {
                // CSRFミドルウェアのエラーはerror.response.data.messageに格納される
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message);
                } else {
                    switchErrorHandling(err.response.data);
                }
            },
        }
    );
    const logoutMutation = useMutation(
        async () => await axios.post(`${process.env.REACT_APP_API_URL}/logout`),
        {
            // ログアウト成功時には、zustandのStateをリセットして、/に遷移
            onSuccess: () => {
                resetEditedTask();
                navigate('/');
            },
            onError: (err: any) => {
                // CSRFミドルウェアのエラーはerror.response.data.messageに格納される
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message);
                } else {
                    switchErrorHandling(err.response.data);
                }
            },
        }
    );
    return { loginMutation, registerMutation, logoutMutation };
}