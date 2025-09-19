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
			(res = await fetch("http://localhost:8080/auth/login")),
				{
					method: "POST",
					headers: { "Content-Type": "application/json" },
					body: JSON.stringify(email, password),
					credentials: "include",
				};
			if (!res.ok) {
				throw new Error("Wrong credentials");
			}

			const data = await res.json();
			dispatch({ type: "login", payload: data.user });
			localStorage.setItem("accessToken", data.token);
		} catch (err) {
			console.log(err);
		}
	}

	function logout() {
		dispatch({ type: "logout" });
	}

	return (
		<AuthContext.Provider value={{ user, isAuthenticated, login, logout }}>
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
