const axios = require("axios");

const API = "http://localhost:8080";

function authHeader(token) {
    return token ? { Authorization: `Bearer ${token}` } : {};
}

async function getAllUsers(token) {
    const res = await axios.get(`${API}/full_user`, { headers: authHeader(token) });
    return res.data;
}

async function getUsersByRole(role, token) {
    const res = await axios.get(`${API}/full_user`, { headers: authHeader(token) });
    return res.data.filter(u => u.user_role === role);
}

async function getAllEvents(token) {
    const res = await axios.get(`${API}/full_event`, { headers: authHeader(token) });
    return res.data;
}

async function getAllGroups(token) {
    const res = await axios.get(`${API}/full_group`, { headers: authHeader(token) });
    return res.data;
}

async function addParticipantsByGroup(eventId, groupId, token) {
    const res = await axios.post(`${API}/add_event_participants_by_group`, {
        event_id: eventId,
        group_id: groupId
    }, { headers: authHeader(token) });
    return res.data;
}

async function getEventById(id, token) {
    const res = await axios.get(`${API}/one_event/${id}`, { headers: authHeader(token) });
    return res.data;
}

async function getParticipantsByEvent(eventId, token) {
    const res = await axios.get(`${API}/participants_by_event/${eventId}`, { headers: authHeader(token) });
    return res.data;
}

async function addEventImage(eventId, urlImage, token) {
    const res = await axios.post(`${API}/new_event_image`, {
        event_id: eventId,
        url_image: urlImage
    }, { headers: authHeader(token) });
    return res.data;
}

async function getEventImagesByEvent(eventId, token) {
    const res = await axios.get(`${API}/event_image_by_event/${eventId}`, { headers: authHeader(token) });
    return res.data;
}

async function addEventResult(eventId, result, token) {
    const res = await axios.post(`${API}/new_event_result`, {
        event_id: eventId,
        result: result
    }, { headers: authHeader(token) });
    return res.data;
}

async function toggleEventStatus(eventId, token) {
    const res = await axios.put(`${API}/event_status/${eventId}`, {}, { headers: authHeader(token) });
    return res.data;
}

async function toggleParticipantStatus(participantId, token) {
    const res = await axios.put(`${API}/toggle_participant_status/${participantId}`, {}, { headers: authHeader(token) });
    return res.data;
}

async function getEventResultsByEvent(eventId, token) {
    const res = await axios.get(`${API}/event_results_by_event/${eventId}`, { headers: authHeader(token) });
    return res.data;
}

async function createGroup(name, token) {
    const res = await axios.post(`${API}/new_group`, { name }, { headers: authHeader(token) });
    return res.data;
}

async function updateGroup(id, name, token) {
    const res = await axios.put(`${API}/update_group/${id}`, { name }, { headers: authHeader(token) });
    return res.data;
}

async function deleteGroup(id, token) {
    const res = await axios.delete(`${API}/delete_group/${id}`, { headers: authHeader(token) });
    return res.data;
}

async function getAllUserGroups(token) {
    const res = await axios.get(`${API}/full_user_group`, { headers: authHeader(token) });
    return res.data;
}

async function createUserGroup(user_id, group_id, token) {
    const res = await axios.post(`${API}/new_user_group`, { user_id, group_id }, { headers: authHeader(token) });
    return res.data;
}

async function updateUserGroup(id, user_id, group_id, token) {
    const res = await axios.put(`${API}/update_user_group/${id}`, { user_id, group_id }, { headers: authHeader(token) });
    return res.data;
}

async function deleteUserGroup(id, token) {
    const res = await axios.delete(`${API}/delete_user_group/${id}`, { headers: authHeader(token) });
    return res.data;
}

async function loginUser(login, password) {
    const res = await axios.post(`${API}/login`, { login, password });
    return res.data;
}

async function registerUser(userData) {
    const res = await axios.post(`${API}/new_user`, userData);
    return res.data;
}

async function addParticipant(eventId, userId, token) {
    const res = await axios.post(
        `${API}/new_event_participant`,
        {
            event_id: eventId,
            user_id: userId
        },
        {
            headers: authHeader(token)
        }
    );

    return res.data;
}

async function getAllParticipants(token) {
    const res = await axios.get(`${API}/full_event_participant`, {
        headers: authHeader(token)
    });

    return Array.isArray(res.data) ? res.data : [];
}

async function deleteParticipant(id, token) {
    const res = await axios.delete(`${API}/delete_event_participant/${id}`, {
        headers: authHeader(token)
    });
    return res.data;
}

async function getUserById(id, token) {
    const res = await axios.get(`${API}/one_user/${id}`, {
        headers: authHeader(token)
    });
    return res.data;
}

async function updateEvent(id, data, token) {
    const res = await axios.put(`${API}/update_event/${id}`, data, { headers: authHeader(token) });
    return res.data;
}
 
async function deleteEvent(id, token) {
    const res = await axios.delete(`${API}/delete_event/${id}`, { headers: authHeader(token) });
    return res.data;
}

module.exports = {
    getAllUsers,
    getUsersByRole,
    getAllEvents,
    getAllGroups,
    addParticipantsByGroup,
    getEventById,
    getParticipantsByEvent,
    addEventImage,
    getEventImagesByEvent,
    addEventResult,
    toggleEventStatus,
    toggleParticipantStatus,
    getEventResultsByEvent,
    createGroup,
    updateGroup,
    deleteGroup,
    getAllUserGroups,
    createUserGroup,
    updateUserGroup,
    deleteUserGroup,
    loginUser,
    registerUser,
    addParticipant,
    getAllParticipants,
    deleteParticipant,
    getUserById,
    updateEvent,
    deleteEvent,
};