import { ArrowRightOnRectangleIcon, ShieldCheckIcon } from "@heroicons/react/24/solid";
import { useMutateAuth  } from "../hooks/useMutateAuth";

export const Todo = () => {
  // Stateの定義
  const { logoutMutation } = useMutateAuth();
  const logout = async () => {
    await logoutMutation.mutateAsync();
  };
  return (
    <div>
      <ArrowRightOnRectangleIcon
        onClick={logout}
        className="h-6 w-6 my-6 cursor-pointer text-blue-500"
      />
    </div>
  )
}
