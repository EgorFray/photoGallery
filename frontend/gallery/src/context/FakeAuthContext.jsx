import { useReducer } from "react";
import { createContext, useContext } from "react";

const AuthContext = createContext();

const initialState = {
	user: null,
	isAuthenticated: false,
};

function reducer(state, action) {
	switch (action.type) {
		case "login":
			return { ...state, user: action.payload, isAuthenticated: true };

		case "logout":
			return { ...state, user: null, isAuthenticated: false };

		default:
			throw new Error("Unknown action");
	}
}

function AuthProvider({ children }) {
	const [{ user, isAuthenticated }, dispatch] = useReducer(reducer, initialState);

	async function login(email, password) {
		try {
			const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/auth/login`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ email, password }),
				credentials: "include",
			});

			if (!res.ok) {
				alert("Wrong email or password");
				throw new Error("Wrong credentials");
			}

			const data = await res.json();
			dispatch({ type: "login", payload: data.user });
			localStorage.setItem("accessToken", data.token);
		} catch (err) {
			throw new Error(err);
		}
	}

	function logout() {
		dispatch({ type: "logout" });
	}

	async function fetchWithAuth(url, options = {}) {
		const headers = {
			...options.headers,
		};

		const accessToken = localStorage.getItem("accessToken");
		if (accessToken) {
			headers["Authorization"] = `Bearer ${accessToken}`;
		}

		let res = await fetch(url, { ...options, headers, credentials: "include" });

		if (res.status == 401) {
			const refreshRes = await fetch(
				`${import.meta.env.VITE_BACKEND_URL}/auth/refresh`,
				{
					method: "POST",
					credentials: "include",
				}
			);
			if (refreshRes.ok) {
				const data = await refreshRes.json();
				const newToken = data.token;
				localStorage.setItem("accessToken", newToken);
				headers["Authorization"] = `Bearer ${newToken}`;
				res = await fetch(url, { ...options, headers, credentials: "include" });
			}
		}

		if (!res.ok) {
			throw new Error(`Request failed: ${res.status}`);
		}
		return res.json();
	}

	return (
		<AuthContext.Provider
			value={{ user, isAuthenticated, login, logout, fetchWithAuth }}
		>
			{children}
		</AuthContext.Provider>
	);
}

function useAuth() {
	const context = useContext(AuthContext);
	if (context === undefined)
		throw new Error("Auth Context was used outside Auth Provider");
	return context;
}

export { AuthProvider, useAuth };
