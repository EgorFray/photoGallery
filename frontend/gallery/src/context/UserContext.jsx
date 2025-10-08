import { createContext, useContext, useState } from "react";

const UserContext = createContext();

function UserProvider({ children }) {
	const [user, setUser] = useState({});
	const [isLoading, setIsLoading] = useState(false);

	async function createUser(newUser) {
		try {
			setIsLoading(true);
			const data = await fetch(`${import.meta.env.BACKEND_URL}/user/create`, {
				method: "POST",
				body: newUser,
			});
			setUser(data);
		} catch {
			alert("There was an error loading data");
		} finally {
			setIsLoading(false);
		}
	}

	return (
		<UserContext.Provider value={{ user, isLoading, createUser }}>
			{children}
		</UserContext.Provider>
	);
}

function useUser() {
	const context = useContext(UserContext);
	if (context === undefined) {
		throw new Error("User Context was used outside User Provider");
	}
	return context;
}

export { UserProvider, useUser };
