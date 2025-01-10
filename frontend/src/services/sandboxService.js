import axios from "axios";

class SandboxService {

    constructor() {
        this.apiClient = axios.create({
            baseURL: "http://localhost:8080",
            headers: {
                "Content-Type": "application/json",
                "Access-Control-Allow-Origin": "*",
            },
        });
    }

    async getAllSandboxes() {
        const response = await this.apiClient.get("/api/sandboxes");
        console.log(response)
        return response.data;
    }

    async deleteSandbox(id) {
        const response = await this.apiClient.delete(`/api/sandboxes/${id}`);
        return response.data;
    }

    async createSandbox() {
        var data = {
            "image_name": "dockware/dev:6.6.8.2",
            "lifetime": 1440
        };
        const response = await this.apiClient.post("/api/sandboxes", sandboxData);
        return response.data;
    }

}

export default new SandboxService();