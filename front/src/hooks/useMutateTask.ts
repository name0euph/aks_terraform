import axios from "axios";
import { useQueryClient, useMutation } from "@tanstack/react-query";
import { Task } from "../types";
import useStore from "../store";
import { useError } from "../hooks/useError";

export const useMutateTask = () => {
    const queryClient = useQueryClient();
    const { switchErrorHandling } = useError();
    const resetEditedTask = useStore((state) => state.resetEditedTask);

    const createTaskMutation = useMutation(
        // taskのid, created_at, updated_atはサーバー側で自動生成されるので、送信しない
        (task: Omit<Task, 'id' | 'created_at' | 'updated_at'>) =>
            axios.post<Task>(`${process.env.REACT_APP_API_URL}/tasks`, task),
        {
            onSuccess: (res) => {
                // タスクの追加が成功したら、tasksのクエリを再取得してキャッシュを更新
                const previousTasks = queryClient.getQueryData<Task[]>(['tasks']);
                if (previousTasks) {
                    // 既存のキャッシュがある場合は、配列の末尾に新しいタスクを追加してキャッシュを更新
                    queryClient.setQueryData(['tasks'], [...previousTasks, res.data]);
                }
                resetEditedTask();
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

    const updateTaskMutation = useMutation(
        // taskのcreated_at, updated_atはサーバー側で自動生成されるので、送信しない
        (task: Omit<Task, 'created_at' | 'updated_at'>) =>
            axios.put<Task>(`${process.env.REACT_APP_API_URL}/tasks/${task.id}`, 
            { 
                title: task.title,
            }
            ),
        {
            onSuccess: (res, variables) => {
                // タスクの更新が成功したら、tasksのクエリを再取得してキャッシュを更新
                const previousTasks = queryClient.getQueryData<Task[]>(['tasks']);
                if (previousTasks) {
                    // 既存のキャッシュがある場合は、更新したIDのタスクに一致するものを新しいタスクに置き換えてキャッシュを更新
                    queryClient.setQueryData<Task[]>(
                        ['tasks'],
                        previousTasks.map((task) =>
                            task.id === variables.id ? res.data : task
                        )
                    );
                }
                resetEditedTask();
            },
            onError: (err: any) => {
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message);
                } else {
                    switchErrorHandling(err.response.data);
                }
            },
        }
    );

    const deleteTaskMutation = useMutation(
        (id: number) =>
            axios.delete<Task>(`${process.env.REACT_APP_API_URL}/tasks/${id}`),
        {
            onSuccess: (_, variables) => {
                // タスクの削除が成功したら、tasksのクエリを再取得してキャッシュを更新
                const previousTasks = queryClient.getQueryData<Task[]>(['tasks']);
                if (previousTasks) {
                    // 既存のキャッシュがある場合は、削除したIDのタスクを除外してキャッシュを更新
                    queryClient.setQueryData<Task[]>(
                        ['tasks'],
                        previousTasks.filter((task) => task.id !== variables)
                    );
                }
            },
            onError: (err: any) => {
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message);
                } else {
                    switchErrorHandling(err.response.data);
                }
            },
        }
    );
    return {
        createTaskMutation,
        updateTaskMutation,
        deleteTaskMutation,
    }
}