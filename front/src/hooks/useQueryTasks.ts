import axios from "axios";
import { useQuery } from "@tanstack/react-query";
import { Task } from "../types";
import { useError } from "../hooks/useError";

export const useQueryTasks = () => {
    const { switchErrorHandling } = useError();
    
    const getTasks = async () => {
        const { data } = await axios.get<Task[]>(
            `${process.env.REACT_APP_API_URL}/tasks`,
            { withCredentials: true }
        );
        return data;
    };

    return useQuery<Task[], Error>({
        // フェッチで取得したデータはクライアントのキャッシュにtasksで格納
        queryKey: ['tasks'],
        queryFn: getTasks,
        staleTime: Infinity,
        onError: (err: any) => {
            // CSRFミドルウェアのエラーはerror.response.data.messageに格納される
            if (err.response.data.message){
                switchErrorHandling(err.response.data.message);
            }
            else {
                switchErrorHandling(err.message);
            }
        }
    });
}