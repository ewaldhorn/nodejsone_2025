// ----------------------------------------------------------------------------
// some variables and constants to HTML elements that are needed
const [incrementElement, countElement] =
  document.querySelectorAll("#increment, #count");
let callbackButton = document.getElementById("callbackButton");
const runtimeInformationElement = document.querySelector("#runtimeInformation");
let populateTableButton = document.getElementById("populateTableButton");
const peopleTableElement = document.getElementById("peopleTable");
const chartElement = document.getElementById("chart-container");

// ----------------------------------------------------------------------------
// refreshes the Go runtime information area
// builds the table information with JS string interpolation
const refreshRuntimeInformation = () => {
  window.getMemoryStats().then((result) => {
    runtimeInformationElement.innerHTML = `<p>Runtime Information (5s delay, click to refresh now)</p>
      <table>
       <tr><td>Allocated </td><td> ${result.allocated} MiB</td></tr>
       <tr><td>Total Allocated </td><td> ${result.totalAllocated} MiB</td><tr>
       <tr><td>Reserved </td><td> ${result.reserved} MiB</td><tr>
       <tr><td>GC Count </td><td> ${result.gc}</td><tr>
      </table>`;
  });
};

// ----------------------------------------------------------------------------
// increment shows interaction between the UI and the backend, creating the
// often-used "click here and see the number increase" example.
document.addEventListener("DOMContentLoaded", () => {
  incrementElement.addEventListener("click", () => {
    window.increment().then((result) => {
      countElement.textContent = result.count;
    });
  });

  document.getElementById("callbackButton").addEventListener("click", () => {
    window.runJS(makeBlueCode).then(() => {});
  });

  populateTableButton.addEventListener("click", () => {
    window.populateTable().then((result) => {
      peopleTableElement.innerHTML = result;
    });
  });

  runtimeInformationElement.addEventListener("click", () => {
    refreshRuntimeInformation();
  });

  // call the Go garbage collector and then refresh the runtime information
  document.getElementById("requestGCButton").addEventListener("click", () => {
    window.gc();
    refreshRuntimeInformation();
  });

  // refresh the runtime information right away
  // and then make it refresh every 5 seconds
  refreshRuntimeInformation();
  setInterval(refreshRuntimeInformation, 5000);

  // make it easy to quit the app via the keyboard shortcut META-Q or CTRL-Q
  document.addEventListener("keydown", function (event) {
    if (event.key == "q" && (event.metaKey || event.ctrlKey)) {
      window.quit();
    }
  });

  renderChart();
});

// ----------------------------------------------------------------------------
const renderChart = () => {
  // Get the canvas element
  const canvas = document.getElementById("chart");
  const ctx = canvas.getContext("2d");

  // Set the canvas dimensions
  canvas.width = 700;
  canvas.height = 500;

  const chartValues = [
    { date: "2013-01-01", rate: 6.55 },
    { date: "2014-01-01", rate: 8.28 },
    { date: "2015-01-01", rate: 9.35 },
    { date: "2016-01-01", rate: 13.85 },
    { date: "2017-01-01", rate: 12.38 },
    { date: "2018-01-01", rate: 11.63 },
    { date: "2019-01-01", rate: 13.43 },
    { date: "2020-01-01", rate: 14.45 },
    { date: "2021-01-01", rate: 15.23 },
    { date: "2022-01-01", rate: 16.12 },
  ];

  // Define the chart dimensions
  const chartWidth = 600;
  const chartHeight = 400;
  const chartX = (canvas.width - chartWidth) / 2;
  const chartY = (canvas.height - chartHeight) / 2;

  // Draw the chart background
  ctx.fillStyle = "#f7f7f7";
  ctx.fillRect(chartX, chartY, chartWidth, chartHeight);

  // Draw the chart grid
  ctx.strokeStyle = "#e0e0e0";
  ctx.lineWidth = 1;
  for (let i = 0; i <= 10; i++) {
    const y = chartY + (chartHeight / 10) * i;
    ctx.beginPath();
    ctx.moveTo(chartX, y);
    ctx.lineTo(chartX + chartWidth, y);
    ctx.stroke();
  }
  for (let i = 0; i <= 10; i++) {
    const x = chartX + (chartWidth / 10) * i;
    ctx.beginPath();
    ctx.moveTo(x, chartY);
    ctx.lineTo(x, chartY + chartHeight);
    ctx.stroke();
  }

  // Draw the values line
  ctx.strokeStyle = "#007bff";
  ctx.lineWidth = 3;
  ctx.beginPath();
  ctx.moveTo(
    chartX,
    chartY +
      chartHeight -
      (chartValues[0].rate /
        Math.max(...chartValues.map((rate) => rate.rate))) *
        chartHeight,
  );
  for (let i = 1; i < chartValues.length; i++) {
    const x = chartX + (chartWidth / (chartValues.length - 1)) * i;
    const y =
      chartY +
      chartHeight -
      (chartValues[i].rate /
        Math.max(...chartValues.map((rate) => rate.rate))) *
        chartHeight;
    ctx.lineTo(x, y);
  }
  ctx.stroke();

  // Add date labels
  ctx.font = "10px Arial";
  ctx.fillStyle = "#ffffff";
  ctx.textAlign = "center";
  for (let i = 0; i < chartValues.length; i++) {
    const x = chartX + (chartWidth / (chartValues.length - 1)) * i;
    const y = chartY + chartHeight + 20;
    ctx.fillText(chartValues[i].date, x, y);
  }
};
