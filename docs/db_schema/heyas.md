# heyas

## Description

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE TABLE `heyas` (
  `id` char(36) NOT NULL,
  `title` char(50) NOT NULL,
  `description` text NOT NULL,
  `creator_id` char(36) NOT NULL,
  `last_editor_id` char(36) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `fk_heyas_creator` (`creator_id`),
  KEY `fk_heyas_last_editor` (`last_editor_id`),
  CONSTRAINT `fk_heyas_creator` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_heyas_last_editor` FOREIGN KEY (`last_editor_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
```

</details>

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| id | char(36) |  | false | [favorites](favorites.md) [hiqidashis](hiqidashis.md) [histories](histories.md) [tsunas](tsunas.md) |  |  |
| title | char(50) |  | false |  |  |  |
| description | text |  | false |  |  |  |
| creator_id | char(36) |  | false |  | [users](users.md) |  |
| last_editor_id | char(36) |  | false |  | [users](users.md) |  |
| created_at | datetime | current_timestamp() | false |  |  |  |
| updated_at | datetime | current_timestamp() | false |  |  |  |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| fk_heyas_creator | FOREIGN KEY | FOREIGN KEY (creator_id) REFERENCES users (id) |
| fk_heyas_last_editor | FOREIGN KEY | FOREIGN KEY (last_editor_id) REFERENCES users (id) |
| PRIMARY | PRIMARY KEY | PRIMARY KEY (id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| fk_heyas_creator | KEY fk_heyas_creator (creator_id) USING BTREE |
| fk_heyas_last_editor | KEY fk_heyas_last_editor (last_editor_id) USING BTREE |
| PRIMARY | PRIMARY KEY (id) USING BTREE |

## Relations

![er](heyas.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
