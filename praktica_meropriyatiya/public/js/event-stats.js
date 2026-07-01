const canvas = document.getElementById('attendanceChart');

const came = Number(canvas.dataset.came);
const absent = Number(canvas.dataset.absent);

new Chart(canvas, {
    type: 'pie',
    data: {
        labels: ['Пришли', 'Отсутствовали'],
        datasets: [{
            data: [came, absent],
            backgroundColor: [
                '#22c55e',
                '#ef4444'
            ],
            borderWidth: 1
        }]
    },
    options: {
        responsive: true,
        plugins: {
            legend: {
                position: 'bottom'
            },
            title: {
                display: true,
                text: 'Посещаемость мероприятия'
            }
        }
    }
});