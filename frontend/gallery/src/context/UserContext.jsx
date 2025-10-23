import { createContext, useContext, useState } from "react";
import { useAuth } from "./FakeAuthContext";

const UserContext = createContext();

function UserProvider({ children }) {
	const { fetchWithAuth } = useAuth();
	const [user, setUser] = useState({});
	const [isLoading, setIsLoading] = useState(false);

	async function createUser(newUser) {
		try {
			setIsLoading(true);
			const data = await fetch(`${import.meta.env.VITE_BACKEND_URL}/user/create`, {
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

	async function updateUser(updatedUser) {
		try {
			setIsLoading(true);
			const data = await fetchWithAuth(
				`${import.meta.env.VITE_BACKEND_URL}/user/update`,
				{
					method: "PATCH",
					body: updatedUser,
				}
			);
			setUser(data);
		} catch {
			alert("There was an error loading data");
		} finally {
			setIsLoading(false);
		}
	}

	return (
		<UserContext.Provider value={{ user, isLoading, createUser, updateUser }}>
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
