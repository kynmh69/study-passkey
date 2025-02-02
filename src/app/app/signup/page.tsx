export default function Signup() {
    return (
        <div className={"flex justify-center"}>
            {/* card layout */}
            <div className={"container mx-auto p-4 bg-white shadow-md rounded-md mt-4"}>
                <h1 className={"font-bold text-2xl "}>Sign Up</h1>
                <form>
                    {/*email*/}
                    <div className={"mb-4"}>
                        <label htmlFor={"email"} className={"block text-sm font-medium text-gray-700"}>
                            Email
                        </label>
                        <input
                            type={"email"}
                            id={"email"}
                            name={"email"}
                            className={"mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"}
                            placeholder={"Email"}/>
                    </div>
                    {/*username*/}
                    <div className={"mb-4"}>
                        <label htmlFor={"username"} className={"block text-sm font-medium text-gray-700"}>
                            Username
                        </label>
                        <input
                            type={"text"}
                            id={"username"}
                            name={"username"}
                            className={"mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"}
                            placeholder={"Username"}/>
                    </div>
                    {/*sign up with passkey*/}
                    <div className={"mb-4"}>
                        <button
                            type={"submit"}
                            className={"w-full bg-violet-700 text-white p-2 rounded-md"}>
                            Sign Up
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
}