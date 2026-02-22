```
db.books.insertOne({
  title: "gRPC",
  author: "Касун Индрасири",
  year: 2020
});
```

```
db.books.insertMany([
  { title: "Clean Code", author: "Robert Martin", year: 2008 },
  { title: "Designing Data-Intensive Applications", author: "Martin Kleppmann", year: 2017 },
  { title: "MongoDB Basics", author: "Jane Doe", year: 2021 }
]);
```

```
db.books.find({ year: { $gte: 2017 } });
```

```
db.books.replaceOne(
  { title: "Go" },
  { title: "Go in Action", author: "William Kennedy", year: 2016 },
  { upsert: false }
);
```

```
db.books.updateOne(
  { title: "Clean Code" },
  { $set: { genre: "Software Engineering" } }
);
```

```
db.books.deleteOne({ year: 2020 });
```

```
db.books.countDocuments()
```

```
db.books.updateMany(
  { rating: null },
  { $set: { rating: 3 } }
)
```
