import { useEffect } from "react";
import { createContext, useContext, useState } from "react";
import { useAuth } from "./FakeAuthContext";

const UserContext = createContext();

function UserProvider({ children }) {
	const { fetchWithAuth } = useAuth();
	const [newUser, setNewUser] = useState({});
	const [curUser, setCurUser] = useState({});
	const [isLoading, setIsLoading] = useState(false);

	async function getCurrentUser(userId) {
		try {
			setIsLoading(true);
			const data = await fetchWithAuth(`${import.meta.env.VITE_BACKEND_URL}/user`, {
				method: "GET",
				body: userId,
			});
			setCurUser(data);
		} catch {
			alert("There was an error loading data");
		} finally {
			setIsLoading(false);
		}
	}
	async function createUser(newUser) {
		try {
			setIsLoading(true);
			const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/user/create`, {
				method: "POST",
				body: newUser,
			});
			const data = await res.json();
			if (!res.ok) {
				throw new Error(data.Error || "Failed to create user");
			}
			setNewUser(data);
		} catch (err) {
			if (err.message.includes("email")) {
				alert("This email has already been used");
			}
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
			setNewUser(data);
		} catch {
			alert("There was an error loading data");
		} finally {
			setIsLoading(false);
		}
	}

	return (
		<UserContext.Provider
			value={{ newUser, curUser, isLoading, createUser, getCurrentUser, updateUser }}
		>
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
