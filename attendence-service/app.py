from flask import Flask, jsonify, request

app = Flask(__name__)

# Dummy attendance data
attendance = [
    {"id": 1, "user": "prudhvi", "status": "Present"},
    {"id": 2, "user": "raj", "status": "Absent"}
]

@app.route('/')
def home():
    return "Welcome to the Attendance Service!"

@app.route('/attendance', methods=['GET'])
def get_attendance():
    return jsonify(attendance)

@app.route('/attendance', methods=['POST'])
def mark_attendance():
    data = request.get_json()
    attendance.append(data)
    return jsonify({"message": "Attendance marked"}), 201

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
