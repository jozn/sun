CREATE
    TRIGGER `ms`.`add_post` AFTER UPDATE
    ON `ms`.`tags`
    FOR EACH ROW BEGIN
	INSERT INTO log_changes (Id2,T) VALUES (New.Id, "tag")
    END

DELIMITER ;